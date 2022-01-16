package main

import (
	"helia/engine"
	"helia/listener"
	"helia/shared"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// configure global tee logging
	shared.InitializeTeeLog(
		printLogger,
	)

	// initialize RNG
	shared.TeeLog("Initializing RNG...")
	rand.Seed(time.Now().UnixNano())

	// initialize game engine
	shared.TeeLog("Initializing engine...")
	engine := engine.HeliaEngine{}
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

	// instantiate http listener
	shared.TeeLog("Initializing HTTP listener...")
	httpListener := &listener.HTTPListener{}
	httpListener.Engine = &engine

	// listen an serve api requests
	shared.TeeLog("Wiring up HTTP handlers...")
	http.HandleFunc("/api/register", httpListener.HandleRegister)
	http.HandleFunc("/api/login", httpListener.HandleLogin)
	http.HandleFunc("/api/shutdown", httpListener.HandleShutdown)

	shared.TeeLog("Helia is running!")
	http.ListenAndServe(":8080", nil)
}

func printLogger(s string, t time.Time) {
	log.Println(s) // t is intentionaly discarded because Println already provides a timestamp
}
