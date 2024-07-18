package repository

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/amirnilofari/library-system/pkg/model"
)

// GetAllBooks retrieves all books from the database
func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	query := "SELECT id, title, author, published_date, isbn, available_copies FROM books;"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query %s: %v", query, err)
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.ISBN, &book.AvailableCopies); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return books, nil
}

// CreateBook inserts a new book into the database
func CreateBook(db *sql.DB, title, author, publishedDate, isbn, availableCopies string) error {
	query := "INSERT INTO books (title, author, published_date, isbn, available_copies) VALUES (?, ?, ?, ?, ?);"
	availableCopiesInt, err := strconv.Atoi(availableCopies)
	if err != nil {
		return err
	}
	_, err = db.Exec(query, title, author, publishedDate, isbn, availableCopiesInt)
	if err != nil {
		log.Printf("Error executing query %s: %v", query, err)
		return err
	}
	return nil
}
