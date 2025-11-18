package services

import (
	"errors"
	"speakbuddy/internal/models"
	"speakbuddy/internal/repository"
	"speakbuddy/pkg/dto/request"
)

type ConsultationService interface {
	BookConsultation(userID uint, therapistID uint, data request.CreateConsultationRequest) (*models.Consultation, error)
	GetMyConsultations(userID uint) ([]models.Consultation, error)
	GetTherapistConsultations(therapistID uint) ([]models.Consultation, error)
}

type consultationService struct {
	repo repository.ConsultationRepository
}

func NewConsultationService(repo repository.ConsultationRepository) ConsultationService {
	return &consultationService{repo}
}

func (s *consultationService) BookConsultation(userID uint, therapistUserID uint, data request.CreateConsultationRequest) (*models.Consultation, error) {
	// Therapist harus punya role "therapist"
	if data.TimeSlot == "" {
		return nil, errors.New("time slot is required")
	}

	consult := &models.Consultation{
		UserID:      userID,
		TherapistUserID: therapistUserID,

		ChildName: data.ChildName,
		ChildAge:  data.ChildAge,
		ChildSex:  data.ChildSex,

		Date:     data.Date,
		TimeSlot: data.TimeSlot,
		IsPaid:   data.IsPaid,
	}

	err := s.repo.Create(consult)
	return consult, err
}

func (s *consultationService) GetMyConsultations(userID uint) ([]models.Consultation, error) {
	return s.repo.FindByUser(userID)
}

func (s *consultationService) GetTherapistConsultations(therapistID uint) ([]models.Consultation, error) {
	return s.repo.FindByTherapist(therapistID)
}
