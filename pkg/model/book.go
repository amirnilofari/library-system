package model

import (
	"database/sql"
	"time"
)

type Book struct {
	ID              int            `json:"id"`
	Title           string         `json:"title"`
	Author          string         `json:"author"`
	PublishedDate   time.Time      `json:"published_date"`
	ISBN            sql.NullString `json:"isbn"`
	AvailableCopies sql.NullInt64  `json:"available_copies"`
}

// func NewBook(title, author, isbn string, publishedDate time.Time, availableCopies int) *Book {
// 	return &Book{
// 		Title:           title,
// 		Author:          author,
// 		PublishedDate:   publishedDate,
// 		ISBN:            isbn,
// 		AvailableCopies: availableCopies,
// 	}
// }

// func (b *Book) IsAvailable() bool {
// 	return b.AvailableCopies > 0
// }
