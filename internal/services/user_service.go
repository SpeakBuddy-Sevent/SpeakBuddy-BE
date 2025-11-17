package services

import (
	"speakbuddy/internal/repository"
)

type UserService interface {
	UpdateName(userID uint, name string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) UpdateName(userID uint, name string) error {
	return s.userRepo.UpdateName(userID, name)
}
