package main

import (
	"helia/engine"
	"helia/listener"
	"net/http"
)

func main() {
	//initialize game engine
	engine := engine.HeliaEngine{}
	engine.Initialize()

	//instantiate socket listener
	socketListener := &listener.SocketListener{}
	socketListener.Engine = &engine
	http.HandleFunc("/ws/connect", socketListener.HandleConnect)

	//start engine
	engine.Start()

	//listen an serve api requests
	http.HandleFunc("/api/register", listener.HandleRegister)
	http.HandleFunc("/api/login", listener.HandleLogin)

	http.ListenAndServe(":8080", nil)
}
