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
	"runtime"
	"time"
)

func main() {
	// use one less core for goroutines than is available
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	// configure global tee logging
	shared.InitializeTeeLog(
		printLogger,
		dbLogger,
	)

	// purge old logs
	shared.TeeLog("Nuking logs from previous boots...")
	err := sql.GetLogService().NukeLogs()

	if err != nil {
		shared.TeeLog(err.Error())
		panic(err)
	}

	// initialize RNG
	shared.TeeLog("Initializing RNG...")
	rand.Seed(time.Now().UnixNano())

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
		time.Sleep(30000 * time.Millisecond)
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
