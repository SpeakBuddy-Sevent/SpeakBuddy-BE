package services

import (
	"speakbuddy/internal/models"
	"speakbuddy/internal/repository"
)

type UserService interface {
	UpdateName(userID uint, name string) error
	GetByID(userID uint) (*models.User, error)
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

func (s *userService) GetByID(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}

