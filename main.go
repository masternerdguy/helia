package main

import (
	"fmt"
	"helia/engine"
	"helia/listener"
	"helia/sql"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// purge old logs
	log.Println("Nuking logs from previous boots...")
	err := sql.GetLogService().NukeLogs()

	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	// initialize RNG
	log.Println("Initializing RNG...")
	rand.Seed(time.Now().UnixNano())

	// brief sleep
	time.Sleep(100 * time.Millisecond)

	// instantiate http listener
	log.Println("Initializing HTTP listener...")
	httpListener := &listener.HTTPListener{}
	httpListener.Initialize()

	// listen for pings early
	go func() {
		log.Println("Hooking early ping listener...")
		http.HandleFunc("/", httpListener.HandlePing)

		http.ListenAndServe(fmt.Sprintf(":%v", httpListener.GetPort()), nil)
	}()

	// run daily downtime jobs
	log.Println("Running downtime jobs...")
	downtimeRunner := engine.DownTimeRunner{}
	downtimeRunner.Initialize()
	downtimeRunner.RunDownTimeTasks()

	// initialize game engine
	log.Println("Initializing engine...")
	engine := engine.HeliaEngine{}
	httpListener.Engine = &engine
	engine.Initialize()

	// instantiate socket listener
	log.Println("Initializing socket listener...")
	socketListener := &listener.SocketListener{}
	socketListener.Initialize()
	socketListener.Engine = &engine

	log.Println("Wiring up socket handlers...")
	http.HandleFunc("/ws/connect", socketListener.HandleConnect)

	// start engine
	log.Println("Starting engine...")
	engine.Start()

	// listen and serve api requests
	log.Println("Wiring up API HTTP handlers...")
	http.HandleFunc("/api/register", httpListener.HandleRegister)
	http.HandleFunc("/api/login", httpListener.HandleLogin)
	http.HandleFunc("/api/shutdown", httpListener.HandleShutdown)

	// give the user a chance to accept the self signed cert
	http.HandleFunc("/dev/accept-cert", httpListener.HandleAcceptCert)

	// up and running!
	log.Println("Helia is running!")

	// don't exit
	for {
		time.Sleep(5000 * time.Millisecond)
	}
}

func printLogger(s string, t time.Time) {
	log.Println(s) // t is intentionally discarded because Println already provides a timestamp
}

func dbLogger(s string, t time.Time) {
	// get log service
	logSvc := sql.GetLogService()

	// write log
	logSvc.WriteLog(s, t)
}
