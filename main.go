package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("Display a specific snippet...%d", id)))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check weather the request is post or not

	if r.Method != "POST" {

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

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
