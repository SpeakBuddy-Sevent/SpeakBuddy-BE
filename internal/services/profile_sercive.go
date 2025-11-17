package services


import (
	// "errors"
	"speakbuddy/internal/models"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/internal/repository"
)


type ProfileService interface {
	CreateOrUpdateProfile(userID uint, req request.CreateProfileRequest) (*models.Profile, error)
	GetProfile(userID uint) (*models.Profile, error)
}


type profileService struct {
	repo repository.ProfileRepository
}


func NewProfileService(repo repository.ProfileRepository) ProfileService {
	return &profileService{repo}
}


func (s *profileService) CreateOrUpdateProfile(userID uint, req request.CreateProfileRequest) (*models.Profile, error) {
	prof, err := s.repo.FindByUserID(userID)
	if err != nil {
	// Buat baru
		newProf := models.Profile{
		UserID: userID,
		Age: req.Age,
		Sex: req.Sex,
		Phone: req.Phone,
		}
	err = s.repo.Create(&newProf)
	return &newProf, err
	}


	// Update
	prof.Age = req.Age
	prof.Sex = req.Sex
	prof.Phone = req.Phone


	err = s.repo.Update(prof)
	return prof, err
}


func (s *profileService) GetProfile(userID uint) (*models.Profile, error) {
	return s.repo.FindByUserID(userID)
}