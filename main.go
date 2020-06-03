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

	//listen an serve api requests
	http.HandleFunc("/api/register", listener.HandleRegister)
	http.HandleFunc("/api/login", listener.HandleLogin)

	http.ListenAndServe(":8080", nil)
}
