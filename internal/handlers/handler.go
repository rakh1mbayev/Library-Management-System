package handlers

import (
	"Library-Management-System/internal/services"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct{}

func (H *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	bookService := services.NewFileBookService()
	bookHandler := NewBookHandler(bookService)
	mux.HandleFunc("POST /book", BookHandler.Add)
}
