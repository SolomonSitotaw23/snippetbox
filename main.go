package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet box"))
}

func main() {
	// initialize a new serve mux
	mux := http.NewServeMux()
	// register home as a function handler for /
	mux.HandleFunc("/", home)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
