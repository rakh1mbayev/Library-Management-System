package main

import (
	book "Library-Management-System/internal/books"
	"log"
	"net/http"
)

func main() {
	initMongoDB()

	http.HandleFunc("/books", book.ListBooks)         // GET all books
	http.HandleFunc("/books/create", book.CreateBook) // POST a new book

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
