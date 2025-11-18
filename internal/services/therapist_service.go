package services

import (
	"speakbuddy/internal/models"
	"speakbuddy/internal/repository"
)

type TherapistService interface {
	GetAllTherapists() ([]models.User, error)
	GetTherapistByID(id uint) (*models.User, error)
}

type therapistService struct {
	userRepo repository.UserRepository
}

func NewTherapistService(userRepo repository.UserRepository) TherapistService {
	return &therapistService{
		userRepo: userRepo,
	}
}

func (s *therapistService) GetAllTherapists() ([]models.User, error) {
	return s.userRepo.FindAllTherapists()
}

func (s *therapistService) GetTherapistByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}
