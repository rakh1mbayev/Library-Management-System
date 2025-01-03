package main

import (
	"Library-Management-System/config"
	"Library-Management-System/internal/handlers"
	"log"
	"net/http"
)

func main() {
	Init()

	handleHTTP := &handlers.Handler{}

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":"+*config.Port, handleHTTP); err != nil {
		log.Fatal(err)
	}
}
