package main

import (
	"fmt"
	"helia/engine"
	"net/http"
)

func main() {
	//initialize game engine
	engine := engine.HeliaEngine{}
	engine.Initialize()

	//start engine
	engine.Start()

	//listen an serve http (temporary block to prevent exit)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
