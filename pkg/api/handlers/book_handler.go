package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/amirnilofari/library-system/pkg/repository"
)

var templ = template.Must(template.ParseGlob("templates/pages/home.html"))

func init() {
	template.Must(templ.ParseGlob("templates/index.html"))
	template.Must(templ.ParseGlob("templates/partials/*.html"))
}

// HomeHandler handles requests to the home page
func HomeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := repository.GetAllBooks(db)
		if err != nil {
			log.Printf("Unable to fetch books: %v", err)
			http.Error(w, "Unable to fetch books", http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "base", map[string]interface{}{
			"Title": "Home",
			"Books": books,
		})
		if err != nil {
			log.Printf("Error rendering books template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// // BooksHandler handles requests to list books
// func BooksHandler(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		books, err := repository.GetAllBooks(db)
// 		if err != nil {
// 			log.Printf("Unable to fetch books: %v", err)
// 			http.Error(w, "Unable to fetch books", http.StatusInternalServerError)
// 			return
// 		}
// 		for _, v := range books {
// 			fmt.Println(v)
// 		}
// 		err = templ.ExecuteTemplate(w, "base", map[string]interface{}{
// 			"Title": "Books",
// 			"Books": books,
// 		})
// 		if err != nil {
// 			log.Printf("Error rendering books template: %v", err)
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		}
// 	}
// }

func CreateBookHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		author := r.FormValue("author")
		publishedDate := r.FormValue("published_date")
		isbn := r.FormValue("isbn")
		availableCopies := r.FormValue("available_copies")

		parsedDate, _ := time.Parse("2006-01-02", publishedDate)
		fmt.Println("unparsed", publishedDate)
		fmt.Println("parsed: ", parsedDate)
		err := repository.CreateBook(db, title, author, isbn, availableCopies, publishedDate)
		if err != nil {
			log.Printf("Unable to create book: %v", err)
			http.Error(w, "Unable to create book", http.StatusInternalServerError)
			return
		}

		// Fetch updated book list
		books, err := repository.GetAllBooks(db)
		if err != nil {
			log.Printf("Unable to fetch books: %v", err)
			http.Error(w, "Unable to fetch books", http.StatusInternalServerError)
			return
		}

		// Respond with updated book list
		err = templ.ExecuteTemplate(w, "book_list", map[string]interface{}{
			"Books": books,
		})
		if err != nil {
			log.Printf("Error rendering book list template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
