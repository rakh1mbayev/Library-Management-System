package models

import "go.mongodb.org/mongo-driver/mongo"

type Book struct {
	ID            int
	Title         string
	Author        string
	ISBN          string
	Publisher     string
	YearPublished int
	Category      string
}

var Client *mongo.Client

// Books is an in-memory storage for Books
var Books = []Book{}
