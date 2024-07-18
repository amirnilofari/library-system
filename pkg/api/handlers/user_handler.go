package handlers

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/amirnilofari/library-system/pkg/repository"
)

func init() {
	template.Must(templ.ParseGlob("templates/index.html"))
	template.Must(templ.ParseGlob("templates/partials/*.html"))
}

// list users
func UsersHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.GetAllUsers(db)
		if err != nil {
			http.Error(w, "Unable to fetch users", http.StatusInternalServerError)
			return
		}
		templ.ExecuteTemplate(w, "base", map[string]interface{}{
			"Title": "Users",
			"Users": users,
		})
	}
}

func NewUserHandler(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "base", map[string]interface{}{
		"Title":   "New User",
		"Content": "user_form",
	})
}

func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		email := r.FormValue("email")

		err := repository.CreateUser(db, firstName, lastName, email)
		if err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}
