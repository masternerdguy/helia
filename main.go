package main

import (
	"fmt"
	"helia/engine"
	"helia/listener"
	"helia/shared"
	"helia/sql"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/debug"
	"runtime/pprof"
	"time"
)

const PROFILE_CPU = true
const PROFILE_HEAP = true
const GC_PERCENT = 500

// current server health loggerPhase
var loggerPhase = shared.PHASE_STARTUP

// whether to blackhole health logging
var dropHealthLogger = false

// get log service
var logSvc = sql.GetLogService()

// program entry point
func main() {
	// configure global tee logging
	shared.InitializeTeeLog(
		printLogger,
		dbLogger,
		healthLogger,
	)

	// purge old logs
	shared.TeeLog("Removing logs older than 7 days...")
	err := sql.GetLogService().NukeLogs()

	if err != nil {
		shared.TeeLog(err.Error())
		panic(err)
	}

	// initialize RNG
	shared.TeeLog("Initializing RNG...")
	rand.Seed(time.Now().UnixNano())

	// start profiling if requested
	if PROFILE_CPU {
		shared.TeeLog("Starting CPU profiling...")
		shared.CpuProfileFile, err = os.Create("cpu.prof")

		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		if err := pprof.StartCPUProfile(shared.CpuProfileFile); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
	}

	if PROFILE_HEAP {
		shared.TeeLog("Enabling shutdown heap profiling...")
		shared.HeapProfileFile, err = os.Create("heap.prof")

		if err != nil {
			log.Fatal("could not create heap profile: ", err)
		}
	}

	// brief sleep
	time.Sleep(100 * time.Millisecond)

	// instantiate http listener
	shared.TeeLog("Initializing HTTP listener...")
	httpListener := &listener.HTTPListener{}
	httpListener.Initialize()

	// listen for pings early
	go func() {
		// hook listener
		shared.TeeLog("Hooking early ping listener...")
		http.HandleFunc("/", httpListener.HandlePing)

		http.ListenAndServe(fmt.Sprintf(":%v", httpListener.GetPort()), nil)
	}()

	// run daily downtime jobs
	shared.TeeLog("Running downtime jobs...")
	downtimeRunner := engine.DownTimeRunner{}
	downtimeRunner.Initialize()
	downtimeRunner.RunDownTimeTasks()

	// initialize shared configuration
	shared.TeeLog("Initializing shared configuration...")
	shared.InitializeConfiguration()

	// initialize game engine
	shared.TeeLog("Initializing engine...")
	engine := engine.HeliaEngine{}
	httpListener.Engine = &engine
	engine.Initialize()

	// instantiate socket listener
	shared.TeeLog("Initializing socket listener...")
	socketListener := &listener.SocketListener{}
	socketListener.Initialize()
	socketListener.Engine = &engine

	shared.TeeLog("Wiring up socket handlers...")
	http.HandleFunc("/ws/connect", socketListener.HandleConnect)

	// start engine
	shared.TeeLog("Starting engine...")
	engine.Start()

	// give the user a chance to accept the self signed cert
	http.HandleFunc("/dev/accept-cert", httpListener.HandleAcceptCert)

	// handle panics that are otherwise unhandled
	shared.TeeLog("Handling main panics...")
	defer func() {
		if r := recover(); r != nil {
			// log error for inspection
			shared.TeeLog(fmt.Sprintf("main panicked: %v", r))

			// include stack trace
			shared.TeeLog(fmt.Sprintf("stacktrace from panic: \n" + string(debug.Stack())))

			// emergency shutdown
			shared.TeeLog("! Emergency shutdown initiated due to main panic!")
			engine.Shutdown()
		}
	}()

	// up and running!
	shared.TeeLog("Helia is running!")

	// stop tee logging to public to prevent sensitive information disclosure
	loggerPhase = shared.PHASE_RUNNING
	dropHealthLogger = true

	// sync health logger
	shared.SetServerHealth(loggerPhase, "")

	// listen and serve api requests
	shared.TeeLog("Wiring up API HTTP handlers...")
	http.HandleFunc("/api/register", httpListener.HandleRegister)
	http.HandleFunc("/api/login", httpListener.HandleLogin)
	http.HandleFunc("/api/forgot", httpListener.HandleForgot)
	http.HandleFunc("/api/reset", httpListener.HandleReset)
	http.HandleFunc("/api/shutdown", httpListener.HandleShutdown)

	// listeners are ready
	shared.TeeLog("Helia is listening!")

	// logger safety check
	if loggerPhase != shared.PHASE_RUNNING || !dropHealthLogger {
		panic("Incorrect tee logger state")
	}

	// watch for shutdown signal to re-enable tee logging to public
	go func(l *listener.HTTPListener) {
		// exit notification
		defer shared.TeeLog("! Entered public shutdown phase - please wait...")

		// loop while not shutting down
		for !httpListener.Engine.IsShuttingDown() {
			// don't peg cpu
			time.Sleep(5 * time.Second)
		}

		// brief sleep
		time.Sleep(1 * time.Second)

		// entered public shutdown
		loggerPhase = shared.PHASE_SHUTDOWN

		// disable health blackholing
		dropHealthLogger = false
	}(httpListener)

	// don't exit
	for {
		// don't peg cpu
		time.Sleep(1 * time.Minute)

		// disable automatic gc
		debug.SetGCPercent(-1)
		time.Sleep(time.Microsecond * 5)

		// invoke garbage collection and return memory to OS
		debug.FreeOSMemory()

		// restore gc settings
		debug.SetGCPercent(GC_PERCENT)
		time.Sleep(time.Microsecond * 5)
	}
}

// logger function to write to the database
func dbLogger(s string, t time.Time) {
	// write log
	logSvc.WriteLog(s, t)
}

// logger function to write to the console
func printLogger(s string, t time.Time) {
	log.Println(s) // t is intentionally discarded because Println already provides a timestamp
}

// logger function to write to public health ping
func healthLogger(s string, t time.Time) {
	// check for blackholing
	if dropHealthLogger {
		return
	}

	// make timestamp
	tx := t.UnixNano() / int64(time.Millisecond)

	// build message
	u := fmt.Sprintf("[%v] %v", tx, s)

	// update health message
	shared.SetServerHealth(loggerPhase, u)
}
