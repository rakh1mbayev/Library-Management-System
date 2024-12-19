package models

import "go.mongodb.org/mongo-driver/mongo"

//type Book struct {
//    ID             int
//    Title          string
//    Author         string
//    ISBN           string
//    Publisher      string
//    YearPublished  int
//    Category       string
//    CopiesAvailable int
//    TotalCopies    int
//}

var Client *mongo.Client

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Books is an in-memory storage for Books
var Books = []Book{}
