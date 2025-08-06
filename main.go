package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	// check if the current request URL path exactly matches "/". as / is a subtree pattern
	// it is not good to route all not found paths to this rout

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// if the path is correct we will display this

	w.Write([]byte("Hello from Snippet box"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func main() {
	// initialize a new serve mux
	mux := http.NewServeMux()
	// register home as a function handler for /
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
