package service

import (
	"database/sql"

	"github.com/amirnilofari/library-system/pkg/model"
	"github.com/amirnilofari/library-system/pkg/repository"
)

type BookService struct {
	DB *sql.DB
}

// Return a list of all books
func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return repository.GetAllBooks(s.DB)
}

// creates a new book
func (s *BookService) CreateBook(title, author, published_date, isbn, available_copies string) error {
	return repository.CreateBook(s.DB, title, author, published_date, isbn, available_copies)
}
