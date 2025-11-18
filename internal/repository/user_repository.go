package repository

import (
	"speakbuddy/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	UpdateName(userID uint, name string) error
	FindAllTherapists() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) UpdateName(userID uint, name string) error {
	return r.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("name", name).Error
}

func (r *userRepository) FindAllTherapists() ([]models.User, error) {
	var users []models.User
	err := r.db.Where("role = ?", "therapist").Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}