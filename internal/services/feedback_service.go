package services

import (
	"speakbuddy/internal/models"
	"speakbuddy/internal/providers"
	"speakbuddy/internal/repository"
)

type FeedbackService interface {
	AnalyzeAndSaveFeedback(sessionID uint, text string) (*models.Feedback, error)
}

type feedbackService struct {
	geminiProvider     providers.GeminiProvider
	feedbackRepository repository.FeedbackRepository
}

func NewFeedbackService(gemini providers.GeminiProvider, repo repository.FeedbackRepository) FeedbackService {
	return &feedbackService{
		geminiProvider:     gemini,
		feedbackRepository: repo,
	}
}

func (s *feedbackService) AnalyzeAndSaveFeedback(sessionID uint, text string) (*models.Feedback, error) {
	result, err := s.geminiProvider.GetFeedbackFromGemini(text)
	if err != nil {
		return nil, err
	}

	feedback := &models.Feedback{
		SessionID: sessionID,
		AIModel:   "gemini-2.5-flash",
		Feedback:  result,
	}

	if err := s.feedbackRepository.Create(feedback); err != nil {
		return nil, err
	}

	return feedback, nil
}
