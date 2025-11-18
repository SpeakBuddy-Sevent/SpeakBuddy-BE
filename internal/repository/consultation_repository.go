package repository

import (
	"speakbuddy/internal/models"

	"gorm.io/gorm"
)

type ConsultationRepository interface {
	Create(data *models.Consultation) error
	FindByUser(userID uint) ([]models.Consultation, error)
	FindByTherapist(therapistID uint) ([]models.Consultation, error)
}

type consultationRepository struct {
	db *gorm.DB
}

func NewConsultationRepository(db *gorm.DB) ConsultationRepository {
	return &consultationRepository{db}
}

func (r *consultationRepository) Create(data *models.Consultation) error {
	return r.db.Create(data).Error
}

func (r *consultationRepository) FindByUser(userID uint) ([]models.Consultation, error) {
	var list []models.Consultation
	err := r.db.Where("user_id = ?", userID).Order("date asc").Find(&list).Error
	return list, err
}

func (r *consultationRepository) FindByTherapist(therapistUserID uint) ([]models.Consultation, error) {
	var list []models.Consultation
	err := r.db.Where("therapist_user_id = ?", therapistUserID).Order("date asc").Find(&list).Error
	return list, err
}

