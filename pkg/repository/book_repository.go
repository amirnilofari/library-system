package repository

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/amirnilofari/library-system/pkg/model"
)

// GetAllBooks retrieves all books from the database
func GetAllBooks(db *sql.DB) ([]model.Book, error) {
	query := "SELECT id, title, author, isbn, available_copies, published_date FROM books;"
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query %s: %v", query, err)
		return nil, err
	}
	defer rows.Close()

	var books []model.Book

	for rows.Next() {
		var book model.Book
		var publishedDate []uint8
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.AvailableCopies, &publishedDate); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		log.Printf("Published Date Type: %T, Value: %v", publishedDate, string(publishedDate))

		if publishedDate != nil {
			parsedDate, err := time.Parse("2006-01-02 15:04:05", string(publishedDate))
			if err != nil {
				log.Printf("Error parsing date: %v", err)
				return nil, err
			}
			book.PublishedDate = sql.NullTime{Time: parsedDate, Valid: true}
		} else {
			book.PublishedDate = sql.NullTime{Valid: false}
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
func CreateBook(db *sql.DB, title, author, isbn, availableCopies, publishedDate string) error {
	query := "INSERT INTO books (title, author, isbn, available_copies, published_date) VALUES (?, ?, ?, ?, ?);"
	availableCopiesInt, err := strconv.Atoi(availableCopies)
	if err != nil {
		return err
	}

	// Ensure the publishedDate is in the correct format
	date, err := time.Parse("2006-01-02", publishedDate)
	if err != nil {
		return err
	}
	formattedDate := date.Format("2006-01-02 15:04:05")

	_, err = db.Exec(query, title, author, isbn, availableCopiesInt, formattedDate)
	if err != nil {
		log.Printf("Error executing query %s: %v", query, err)
		return err
	}
	return nil
}
