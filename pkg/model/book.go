package model

import "time"

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublishedDate   time.Time `json:"published_date"`
	ISBN            string    `json:"isbn"`
	AvailableCopies int       `json:"available_copies"`
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
