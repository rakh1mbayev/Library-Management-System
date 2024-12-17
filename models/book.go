package models

type Book struct {
    ID             int
    Title          string
    Author         string
    ISBN           string
    Publisher      string
    YearPublished  int
    Category       string
    CopiesAvailable int
    TotalCopies    int
}