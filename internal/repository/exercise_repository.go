package repository

import (
	"speakbuddy/internal/models"

	"gorm.io/gorm"
)

// ReadingExerciseTemplateRepository untuk manage ReadingExerciseTemplate
type ReadingExerciseTemplateRepository interface {
	FindByID(id uint) (*models.ReadingExerciseTemplate, error)
	FindByLevel(level string) ([]models.ReadingExerciseTemplate, error)
}

type readingExerciseTemplateRepository struct {
	db *gorm.DB
}

func NewReadingExerciseTemplateRepository(db *gorm.DB) ReadingExerciseTemplateRepository {
	return &readingExerciseTemplateRepository{db}
}

func (r *readingExerciseTemplateRepository) FindByID(id uint) (*models.ReadingExerciseTemplate, error) {
	var exercise models.ReadingExerciseTemplate
	err := r.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("item_number ASC")
	}).First(&exercise, id).Error
	return &exercise, err
}

func (r *readingExerciseTemplateRepository) FindByLevel(level string) ([]models.ReadingExerciseTemplate, error) {
	var exercises []models.ReadingExerciseTemplate
	err := r.db.Where("level = ?", level).Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("item_number ASC")
	}).Find(&exercises).Error
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
