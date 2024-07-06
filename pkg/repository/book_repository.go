package repository

import (
	"database/sql"

	"github.com/amirnilofari/library-system/pkg/model"
)

// Retrieves all books from the database
func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	rows, err := db.Query("SELECT id, title, author, published_date, isbn, available_copies FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.AvailableCopies); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// Inserts a new book into the database
func CreateBook(db *sql.DB, title, author string) error {
	_, err := db.Exec("INSERT INTO books (title, author) VALUES (?,?)", title, author)
	return err
}
