package service

import (
	"database/sql"

	"github.com/amirnilofari/library-system/pkg/model"
	"github.com/amirnilofari/library-system/pkg/repository"
)

type UserService struct {
	DB *sql.DB
}

// Returns a list of all users
func (s *UserService) GetUsers() ([]model.User, error) {
	return repository.GetAllUsers(s.DB)
}

func (s *UserService) CreateUser(firstName, lastName, email string) error {
	return repository.CreateUser(s.DB, firstName, lastName, email)
}
