package repository

import (
	"speakbuddy/internal/models"

	"gorm.io/gorm"
)

// ReadingExerciseRepository untuk manage ReadingExercise
type ReadingExerciseRepository interface {
	Create(exercise *models.ReadingExercise) error
	FindByID(id uint) (*models.ReadingExercise, error)
	FindByUserID(userID uint) ([]models.ReadingExercise, error)
}

type readingExerciseRepository struct {
	db *gorm.DB
}

func NewReadingExerciseRepository(db *gorm.DB) ReadingExerciseRepository {
	return &readingExerciseRepository{db}
}

func (r *readingExerciseRepository) Create(exercise *models.ReadingExercise) error {
	return r.db.Create(exercise).Error
}

func (r *readingExerciseRepository) FindByID(id uint) (*models.ReadingExercise, error) {
	var exercise models.ReadingExercise
	err := r.db.Preload("Items").First(&exercise, id).Error
	return &exercise, err
}

func (r *readingExerciseRepository) FindByUserID(userID uint) ([]models.ReadingExercise, error) {
	var exercises []models.ReadingExercise
	err := r.db.Where("user_id = ?", userID).Preload("Items").Find(&exercises).Error
	return exercises, err
}

// ExerciseItemRepository untuk manage ExerciseItem
type ExerciseItemRepository interface {
	FindByID(id uint) (*models.ExerciseItem, error)
	FindByExerciseID(exerciseID uint) ([]models.ExerciseItem, error)
	Create(item *models.ExerciseItem) error
}

type exerciseItemRepository struct {
	db *gorm.DB
}

func NewExerciseItemRepository(db *gorm.DB) ExerciseItemRepository {
	return &exerciseItemRepository{db}
}

func (r *exerciseItemRepository) FindByID(id uint) (*models.ExerciseItem, error) {
	var item models.ExerciseItem
	err := r.db.Preload("Attempts").First(&item, id).Error
	return &item, err
}

func (r *exerciseItemRepository) FindByExerciseID(exerciseID uint) ([]models.ExerciseItem, error) {
	var items []models.ExerciseItem
	err := r.db.Where("exercise_id = ?", exerciseID).Order("item_number").Preload("Attempts").Find(&items).Error
	return items, err
}

func (r *exerciseItemRepository) Create(item *models.ExerciseItem) error {
	return r.db.Create(item).Error
}

// ExerciseAttemptRepository untuk manage ExerciseAttempt
type ExerciseAttemptRepository interface {
	Create(attempt *models.ExerciseAttempt) error
	FindByItemID(itemID uint) ([]models.ExerciseAttempt, error)
}

type exerciseAttemptRepository struct {
	db *gorm.DB
}

func NewExerciseAttemptRepository(db *gorm.DB) ExerciseAttemptRepository {
	return &exerciseAttemptRepository{db}
}

func (r *exerciseAttemptRepository) Create(attempt *models.ExerciseAttempt) error {
	return r.db.Create(attempt).Error
}

func (r *exerciseAttemptRepository) FindByItemID(itemID uint) ([]models.ExerciseAttempt, error) {
	var attempts []models.ExerciseAttempt
	err := r.db.Where("item_id = ?", itemID).Order("created_at DESC").Find(&attempts).Error
	return attempts, err
}
