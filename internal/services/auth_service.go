package services

import (
	"errors"
	"speakbuddy/internal/models"
	"speakbuddy/internal/repository"
	"speakbuddy/pkg/utils"
)

type AuthService interface {
	Register(name, email, password string) (*models.User, error)
	Login(email, password string) (*models.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{userRepo: repo}
}

func (s *authService) Register(name, email, password string) (*models.User, error) {
	hash, _ := utils.HashPassword(password)

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hash,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *authService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	if !utils.CheckPassword(user.PasswordHash, password) {
		return nil, errors.New("password salah")
	}

	return user, nil
}
