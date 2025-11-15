package repository

import (
	"speakbuddy/internal/models"

	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(session *models.Session) error
	FindByID(id uint) (*models.Session, error)
	Update(session *models.Session) error
	FindByUserID(userID uint) ([]models.Session, error)
	SaveRecording(recording *models.SessionRecording) error
	GetRecordingsBySessionID(sessionID uint) ([]models.SessionRecording, error)
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db}
}

func (r *sessionRepository) Create(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) FindByID(id uint) (*models.Session, error) {
	var session models.Session
	err := r.db.Preload("Recordings").Where("id = ?", id).First(&session).Error
	return &session, err
}

func (r *sessionRepository) Update(session *models.Session) error {
	return r.db.Save(session).Error
}

func (r *sessionRepository) FindByUserID(userID uint) ([]models.Session, error) {
	var sessions []models.Session
	err := r.db.Preload("Recordings").Where("user_id = ?", userID).Find(&sessions).Error
	return sessions, err
}

func (r *sessionRepository) SaveRecording(recording *models.SessionRecording) error {
	return r.db.Create(recording).Error
}

func (r *sessionRepository) GetRecordingsBySessionID(sessionID uint) ([]models.SessionRecording, error) {
	var recordings []models.SessionRecording
	err := r.db.Where("session_id = ?", sessionID).Find(&recordings).Error
	return recordings, err
}
