package model

import "time"

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	NationalCode string    `json:"national_code"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}
