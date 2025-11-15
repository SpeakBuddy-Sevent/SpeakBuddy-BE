package services

import (
	"speakbuddy/internal/models"
	"speakbuddy/internal/providers"
	"speakbuddy/internal/repository"
	"speakbuddy/pkg/utils"
)

type ExerciseService interface {
	TranscribeAndAnalyzeAttempt(userID uint, itemID uint, audioBytes []byte) (*models.ExerciseAttempt, error)
}

type exerciseService struct {
	googleSpeech   *providers.SpeechToTextProvider
	geminiProvider providers.GeminiProvider
	attemptRepo    repository.ExerciseAttemptRepository
	itemRepo       repository.ExerciseItemRepository
}

func NewExerciseService(
	googleSpeech *providers.SpeechToTextProvider,
	gemini providers.GeminiProvider,
	attemptRepo repository.ExerciseAttemptRepository,
	itemRepo repository.ExerciseItemRepository,
) ExerciseService {
	return &exerciseService{
		googleSpeech:   googleSpeech,
		geminiProvider: gemini,
		attemptRepo:    attemptRepo,
		itemRepo:       itemRepo,
	}
}

// TranscribeAndAnalyzeAttempt - transcribe audio + analyze dengan Gemini
func (s *exerciseService) TranscribeAndAnalyzeAttempt(userID uint, itemID uint, audioBytes []byte) (*models.ExerciseAttempt, error) {
	// Get target text dari exercise item
	item, err := s.itemRepo.FindByID(itemID)
	if err != nil {
		return nil, err
	}

	// Step 1: Transcribe audio menggunakan Google Speech-to-Text
	transcribedText, err := s.googleSpeech.TranscribeAudio(audioBytes)
	if err != nil {
		return nil, err
	}

	// Step 2: Analisis dengan Gemini
	feedback, err := s.geminiProvider.GetFeedbackFromGemini(item.TargetText, transcribedText)
	if err != nil {
		return nil, err
	}

	// Step 3: Hitung accuracy (similarity antara target dan transcribed)
	accuracy := utils.CalculateAccuracy(item.TargetText, transcribedText)

	// Step 4: Simpan attempt hasil
	attempt := &models.ExerciseAttempt{
		UserID:          userID,
		ItemID:          itemID,
		TranscribedText: transcribedText,
		AIFeedback:      feedback,
		AIModel:         "gemini-2.0-flash",
		Accuracy:        accuracy,
	}

	if err := s.attemptRepo.Create(attempt); err != nil {
		return nil, err
	}

	return attempt, nil
}
