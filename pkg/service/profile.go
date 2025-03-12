package service

import (
	todo "to-do-list"
	"to-do-list/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}

func (p *ProfileService) GetUserProfile(id int) (todo.UserProfile, error) {
	return p.repo.GetUserProfile(id)
}

func (h *ProfileService) UpdateUserProfile(userProfile todo.UserProfile, userID int) error {
	return h.repo.UpdateUserProfile(userProfile, userID)
}
