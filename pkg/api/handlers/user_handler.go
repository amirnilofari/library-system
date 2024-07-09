package handlers

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/amirnilofari/library-system/pkg/repository"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

// list users
func UserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.GetAllUsers(db)
		if err != nil {
			http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
			return
		}
		templ.ExecuteTemplate(w, "users.html", users)
	}
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "user_form.html", nil)
}

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		email := r.FormValue("email")
		// Add other form fields

		err := repository.CreateUser(db, firstName, lastName, email)
		if err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}
