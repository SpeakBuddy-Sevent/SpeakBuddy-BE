package repository

import (
	"speakbuddy/config"
	"speakbuddy/internal/models"
)

type FeedbackRepository interface {
	Create(feedback *models.Feedback) error
}

type feedbackRepository struct{}

func NewFeedbackRepository() FeedbackRepository {
	return &feedbackRepository{}
}

func (r *feedbackRepository) Create(feedback *models.Feedback) error {
	return config.DB.Create(feedback).Error
}
