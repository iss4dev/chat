package service

import (
	todo "to-do-list"
	"to-do-list/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Profile interface {
	GetUserProfile(id int) (todo.UserProfile, error)
	UpdateUserProfile(userProfile todo.UserProfile, userID int) error
}

type Service struct {
	Authorization
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Profile:       NewProfileService(repos.Profile),
	}
}
