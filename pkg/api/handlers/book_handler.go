package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"github.com/amirnilofari/library-system/pkg/repository"
)

//var templ = template.Must(template.ParseGlob("templates/*.html"))

func init() {
	templ = template.Must(template.New("base").ParseGlob("templates/layouts/*.html"))
	templ = template.Must(templ.ParseGlob("templates/partials/*.html"))
	templ = template.Must(templ.ParseGlob("templates/pages/*.html"))
}

// HomeHandler handles requests to the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "base", map[string]interface{}{
		"Title": "Home",
	})
	if err != nil {
		log.Printf("Error rendering home template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// BooksHandler handles requests to list books
func BooksHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := repository.GetAllBooks(db)
		if err != nil {
			log.Printf("Unable to fetch books: %v", err)
			http.Error(w, "Unable to fetch books", http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "base", map[string]interface{}{
			"Title": "Books",
			"Books": books,
		})
		if err != nil {
			log.Printf("Error rendering books template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// Show the new book form
func NewBookHandler(w http.ResponseWriter, r *http.Request) {
	err := templ.ExecuteTemplate(w, "base", map[string]interface{}{
		"Title":   "New Book",
		"Content": "book_form",
	})
	if err != nil {
		log.Printf("Error rendering new book form template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Create a new book
func CreateBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		author := r.FormValue("author")

		err := repository.CreateBook(db, title, author)
		if err != nil {
			log.Printf("Unable to create book: %v", err)
			http.Error(w, "Unable to create book", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/books", http.StatusSeeOther)
	}
}
