package models

import "time"

type User struct {
    ID             int
    FirstName      string
    LastName       string
    Email          string
    PhoneNumber    string
    MembershipType string
    CreatedAt      time.Time
    UpdatedAt      time.Time
}