package main

import (
	"helia/engine"
	"helia/listener"
	"log"
	"net/http"
)

func main() {
	// initialize game engine
	log.Println("Initializing engine...")
	engine := engine.HeliaEngine{}
	engine.Initialize()

	// instantiate socket listener
	log.Println("Initializing socket listener...")
	socketListener := &listener.SocketListener{}
	socketListener.Engine = &engine

	log.Println("Wiring up socket handlers...")
	http.HandleFunc("/ws/connect", socketListener.HandleConnect)

	// start engine
	log.Println("Starting engine...")
	engine.Start()

	// instantiate http listener
	log.Println("Initializing HTTP listener...")
	httpListener := &listener.HTTPListener{}
	httpListener.Engine = &engine

	// listen an serve api requests
	log.Println("Wiring up HTTP handlers...")
	http.HandleFunc("/api/register", httpListener.HandleRegister)
	http.HandleFunc("/api/login", httpListener.HandleLogin)
	http.HandleFunc("/api/shutdown", httpListener.HandleShutdown)

	log.Println("Helia is running!")
	http.ListenAndServe(":8080", nil)
}
