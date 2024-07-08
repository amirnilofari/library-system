package repository

import (
	"database/sql"

	"github.com/amirnilofari/library-system/pkg/model"
)

// retrieves all users from the database
func GetAllUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query("SELECT id, first_name, last_name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(db *sql.DB, firstName, lastName, email string) error {
	_, err := db.Exec("INSERT INTO users (first_name, last_name, email) VALUES (?, ?, ?)", firstName, lastName, email)
	return err
}
