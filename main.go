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

	//start engine
	//engine.Start()

	//listen an serve http
	http.HandleFunc("/api/register", listener.HandleRegister)
	http.HandleFunc("/api/login", listener.HandleLogin)
	http.HandleFunc("/ws/connect", listener.Echo)

	http.ListenAndServe(":8080", nil)
}
