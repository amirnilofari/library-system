package handlers

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/amirnilofari/library-system/pkg/repository"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

// BooksHandler handles requests to list books

func BooksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := repository.GetAllBooks(db)
		if err != nil {
			http.Error(w, "Unable to fetch books", http.StatusInternalServerError)
			return
		}
		templ.ExecuteTemplate(w, "books.html", books)
	}
}

// Show the new book form
func NewBookHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "book_form.html", nil)
}

// Create a new book
func CreateBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		author := r.FormValue("author")
		// Add other form fields as necessary

		err := repository.CreateBook(db, title, author)
		if err != nil {
			http.Error(w, "Unable to create book", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}
