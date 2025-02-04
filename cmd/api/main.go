package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	addr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var cfg config

	// Create a new ServeMux and register the home function as the handler for the "/" URL pattern
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.Parse()

	// Create custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Add log and start the server with the servemux as the root handler
	infoLog.Printf("Starting server on %s", cfg.addr)
	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
