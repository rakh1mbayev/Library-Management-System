package handlers

import (
	"Library-Management-System/internal/dal"
	"Library-Management-System/internal/models"
	"Library-Management-System/internal/services"
	"encoding/json"
	"net/http"
	"strings"
)

type BookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (b *BookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		dal.SendError(http.StatusBadRequest, "Couldn't decode body", err, w)
		return
	}

	if err := b.service.Create(&newBook); err != nil {
		dal.SendError(http.StatusInternalServerError, "Couldn't create book", err, w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (b *BookHandler) Get(w http.ResponseWriter, r *http.Request) {
	books, err := b.service.Get()

	if err != nil {
		dal.SendError(http.StatusInternalServerError, "Couldn't get books", err, w)
		return
	}
	w.Header().Set("Content-type", "application/json")
	if err = json.NewEncoder(w).Encode(books); err != nil {
		return
	}
}

func (b *BookHandler) GetById(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if book, err := b.service.GetById(id); err != nil {
		dal.SendError(http.StatusInternalServerError, "Couldn't get book", err, w)
		return
	}
	// output book
	w.WriteHeader(http.StatusOK)
}

func (b *BookHandler) Update(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	id := ""

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		dal.SendError(http.StatusBadRequest, "Couldn't decode body", err, w)
		return
	}

	if err := b.service.Update(id, &newBook); err != nil {
		dal.SendError(http.StatusInternalServerError, "Couldn't update book", err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (b *BookHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := ""

	if err := b.service.Delete(id); err != nil {
		dal.SendError(http.StatusInternalServerError, "Couldn't delete book", err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
