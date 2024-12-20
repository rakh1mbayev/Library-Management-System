package main

import (
	"Library-Management-System/config"
	book "Library-Management-System/internal/books"
	"log"
	"net/http"
)

func main() {
	// Initialize MongoDB connection
	initMongoDB()

	// Register route handlers
	http.HandleFunc("/books", book.ListBooks)
	http.HandleFunc("/books/id", book.ListBookByID)       // GET a book by ID
	http.HandleFunc("/books/create", book.CreateBook)     // POST a new book
	http.HandleFunc("/books/update", book.UpdateBookById) // PUT to update a book by ID
	http.HandleFunc("/books/delete", book.DeleteBookById) // DELETE a book by ID

	// Start the HTTP server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":"+*config.Port, nil); err != nil {
		log.Fatal(err)
	}
}
