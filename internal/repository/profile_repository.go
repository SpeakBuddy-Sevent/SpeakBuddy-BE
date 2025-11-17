package repository


import "gorm.io/gorm"
import "speakbuddy/internal/models"


type ProfileRepository interface {
	Create(profile *models.Profile) error
	FindByUserID(userID uint) (*models.Profile, error)
	Update(profile *models.Profile) error
}


type profileRepository struct {
	db *gorm.DB
}


func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}


func (r *profileRepository) Create(profile *models.Profile) error {
	return r.db.Create(profile).Error
}


func (r *profileRepository) FindByUserID(userID uint) (*models.Profile, error) {
	var p models.Profile
	err := r.db.Where("user_id = ?", userID).First(&p).Error
	return &p, err
}


func (r *profileRepository) Update(profile *models.Profile) error {
	return r.db.Save(profile).Error
}