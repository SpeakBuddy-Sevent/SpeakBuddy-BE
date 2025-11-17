package repository

import "gorm.io/gorm"
import "speakbuddy/internal/models"

type DataAnakRepository interface {
	Create(data *models.DataAnak) error
	FindByUserID(userID uint) (*models.DataAnak, error)
	Update(data *models.DataAnak) error
}


type dataAnakRepository struct {
	db *gorm.DB
}


func NewDataAnakRepository(db *gorm.DB) DataAnakRepository {
	return &dataAnakRepository{db}
}


func (r *dataAnakRepository) Create(data *models.DataAnak) error {
	return r.db.Create(data).Error
}


func (r *dataAnakRepository) FindByUserID(userID uint) (*models.DataAnak, error) {
	var d models.DataAnak
	err := r.db.Where("user_id = ?", userID).First(&d).Error
	return &d, err
}


func (r *dataAnakRepository) Update(data *models.DataAnak) error {
	return r.db.Save(data).Error
}