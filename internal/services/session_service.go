package services

import (
	"speakbuddy/internal/models"
	"speakbuddy/internal/providers"
	"speakbuddy/internal/repository"
)

type SessionService interface {
	CreateSession(userID uint, targetText string) (*models.Session, error)
	TranscribeAndAnalyze(sessionID uint, audioBytes []byte) (*models.SessionRecording, error)
	GetSessionWithRecordings(sessionID uint) (*models.Session, error)
}

type sessionService struct {
	googleSpeech *providers.SpeechToTextProvider
	gemini       providers.GeminiProvider
	sessionRepo  repository.SessionRepository
}

func NewSessionService(
	googleSpeech *providers.SpeechToTextProvider,
	gemini providers.GeminiProvider,
	sessionRepo repository.SessionRepository,
) SessionService {
	return &sessionService{
		googleSpeech: googleSpeech,
		gemini:       gemini,
		sessionRepo:  sessionRepo,
	}
}

// CreateSession - user memulai latihan membaca baru
func (s *sessionService) CreateSession(userID uint, targetText string) (*models.Session, error) {
	session := &models.Session{
		UserID:     userID,
		TargetText: targetText,
	}

	err := s.sessionRepo.Create(session)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *sessionService) TranscribeAndAnalyze(sessionID uint, audioBytes []byte) (*models.SessionRecording, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, err
	}

	// Transcribe audio dengan Google Speech-to-Text
	transcribedText, err := s.googleSpeech.TranscribeAudio(audioBytes)
	if err != nil {
		return nil, err
	}

	// Analisis dengan Gemini
	feedback, err := s.gemini.GetFeedbackFromGemini(session.TargetText, transcribedText)
	if err != nil {
		return nil, err
	}

	recording := &models.SessionRecording{
		SessionID:       sessionID,
		TranscribedText: transcribedText,
		AIFeedback:      feedback,
		AIModel:         "gemini-2.0-flash",
		Accuracy:        0.0, // TODO: implement accuracy calculation
	}

	err = s.sessionRepo.SaveRecording(recording)
	if err != nil {
		return nil, err
	}

	return recording, nil
}

func (s *sessionService) GetSessionWithRecordings(sessionID uint) (*models.Session, error) {
	session, err := s.sessionRepo.FindByID(sessionID)
	if err != nil {
		return nil, err
	}

	return session, nil
}
