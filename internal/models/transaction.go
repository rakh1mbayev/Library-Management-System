package models

import "time"

type Transaction struct {
	TransactionID int
	BookID        int
	UserID        int
	IssueDate     time.Time
	DueDate       time.Time
	ReturnDate    time.Time
	FineAmount    float64
}
	