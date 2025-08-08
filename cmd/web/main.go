package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/solomonsitotaw23/snippetbox/cmd/middleware"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//accepting port / address where the server will run from a command line arg -addr=PORT
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

	// leveled logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", middleware.Neuter(fileServer)))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	//since the http.Listen and serve uses the default logger we should change it to our custom logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
