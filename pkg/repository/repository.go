package repository

import (
	todo "to-do-list"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username string) (todo.User, error)
}

type Profile interface {
	GetUserProfile(id int) (todo.UserProfile, error)
	UpdateUserProfile(userProfile todo.UserProfile, userID int) error
}

type Repository struct {
	Authorization
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Profile:       NewProfilePostgres(db),
	}
}
