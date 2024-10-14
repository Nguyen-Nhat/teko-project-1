// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"
)

type Author struct {
	ID        int32
	Fullname  string
	Dob       sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	ID                int32
	Title             string
	PublishedYear     int32
	AvailableQuantity int32
	BorrowQuantity    int32
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type BookAuthor struct {
	ID       int32
	BookID   int32
	AuthorID int32
}

type BookGenre struct {
	ID      int32
	BookID  int32
	GenreID int32
}

type BorrowBook struct {
	ID         int32
	StudentID  int32
	BorrowDate time.Time
	ReturnDate time.Time
	IsReturn   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type BorrowBookDetail struct {
	ID           int32
	BorrowBookID int32
	BookID       int32
	Quantity     int32
}

type Genre struct {
	ID        int32
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
