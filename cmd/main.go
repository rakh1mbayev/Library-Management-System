package cmd

import (
	"Library-Management-System/config"
	handler "Library-Management-System/internal/books"
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	http.HandleFunc("/", handler.Handler)

	err := http.ListenAndServe(":"+*config.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
