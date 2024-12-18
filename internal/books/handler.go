package books

import (
	model "Library-Management-System/models"
	"encoding/json"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		ListBooks(w)
	} else if r.Method == "POST" {
		CreateBook(w, r)
	} else {
		http.Error(w, "Only GET and POST methods are supported.", http.StatusMethodNotAllowed)
	}
}
func ListBooks(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(model.Books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Add a new book
	var newBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	model.Books = append(model.Books, newBook)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Book added successfully",
	})
}
