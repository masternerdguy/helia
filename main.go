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
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"time"
)

var profileCpu = true

func main() {
	// configure global tee logging
	shared.InitializeTeeLog(
		printLogger,
		dbLogger,
	)

	// purge old logs
	shared.TeeLog("Removing logs older than 48 hours...")
	err := sql.GetLogService().NukeLogs()

	if err != nil {
		shared.TeeLog(err.Error())
		panic(err)
	}

	// initialize RNG
	shared.TeeLog("Initializing RNG...")
	rand.Seed(time.Now().UnixNano())

	// start profiling if requested
	if profileCpu {
		shared.TeeLog("Starting CPU profiling...")
		shared.CpuProfileFile, err = os.Create("cpu.prof")

		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}

		if err := pprof.StartCPUProfile(shared.CpuProfileFile); err != nil {
			log.Fatal("could not start CPU profile: ", err)
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
		shared.TeeLog("Hooking early ping listener...")
		http.HandleFunc("/", httpListener.HandlePing)

		http.ListenAndServe(fmt.Sprintf(":%v", httpListener.GetPort()), nil)
	}()

	// check whether to use Azure hacks
	if httpListener.UseAzureHacks() {
		shared.TeeLog("Enabling Azure Hacks!!!")

		/* BEGIN AZURE APP SERVICE PERFORMANCE WORKAROUNDS */

		// disable automatic garbage collection :activex: :roach party:
		debug.SetGCPercent(-1)

		// polling-based garbage collection
		go func() {
			/*
			 * This is a workaround to make Helia more cpu efficient when running as a docker container
			 * within an Azure app service. Based on profiling, there are memory allocation issues -
			 * most likely due to heavy iteration over maps. These are a big deal to fix, and I don't
			 * have time right now, so this will act as a bandaid fix until then. Helia actually uses
			 * very little memory, so we can defer running the garbage collector significantly until
			 * traffic gets high enough to overwhelm this hack. Telemetry is logged to help keep an
			 * eye on this known future problem. The long term solution is to convert frequently
			 * iterated maps to slices, which has implications for other things like searching them.
			 *
			 * Note that these garbage collection issues don't occur on any other system that I've
			 * tested, which really shows how weak Azure app services actually are - the other systems
			 * are fast enough that this simply doesn't become a problem and go can manage its own
			 * garbage collection timing like its designed to.
			 */

			gcRuns := 0

			for {
				// throttle rate
				time.Sleep(200 * time.Millisecond)

				// get memory allocation
				var m runtime.MemStats
				runtime.ReadMemStats(&m)

				// convert to megabytes
				commitedMb := 0.000001 * float64(m.Alloc)

				// disgusting... :hug: :party parrot:
				if commitedMb > 5120 {
					// increment gc run counter
					gcRuns++

					// invoke garbage collection
					runtime.GC()

					if gcRuns > 1000 {
						// log memory usage
						shared.TeeLog(fmt.Sprintf("<MEMORY COMMITED> %v [%v gc runs since last log]", commitedMb, gcRuns))

						// reset counter
						gcRuns = 0
					}
				}
			}
		}()

		/* END AZURE APP SERVICE PERFORMANCE WORKAROUNDS */
	}

	// run daily downtime jobs
	shared.TeeLog("Running downtime jobs...")
	downtimeRunner := engine.DownTimeRunner{}
	downtimeRunner.Initialize()
	downtimeRunner.RunDownTimeTasks()

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

	// listen and serve api requests
	shared.TeeLog("Wiring up API HTTP handlers...")
	http.HandleFunc("/api/register", httpListener.HandleRegister)
	http.HandleFunc("/api/login", httpListener.HandleLogin)
	http.HandleFunc("/api/shutdown", httpListener.HandleShutdown)

	// give the user a chance to accept the self signed cert
	http.HandleFunc("/dev/accept-cert", httpListener.HandleAcceptCert)

	// up and running!
	shared.TeeLog("Helia is running!")

	// don't exit
	for {
		// don't peg cpu
		time.Sleep(30 * time.Minute)

		// log goroutine count
		shared.TeeLog(fmt.Sprintf("<TOTAL GOROUTINES> %v", runtime.NumGoroutine()))
	}
}

func printLogger(s string, t time.Time) {
	log.Println(s) // t is intentionally discarded because Println already provides a timestamp
}

// get log service
var logSvc = sql.GetLogService()

func dbLogger(s string, t time.Time) {
	// write log
	logSvc.WriteLog(s, t)
}
