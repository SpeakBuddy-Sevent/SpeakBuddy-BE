package models

import "time"

type Session struct {
	ID         uint               `gorm:"primaryKey" json:"id"`
	UserID     uint               `json:"user_id"`
	TargetText string             `json:"target_text"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Recordings []SessionRecording `gorm:"foreignKey:SessionID" json:"recordings,omitempty"`
}

type SessionRecording struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	SessionID       uint      `json:"session_id"`
	TranscribedText string    `json:"transcribed_text"`
	AIFeedback      string    `gorm:"type:text" json:"feedback"`
	AIModel         string    `json:"ai_model"`
	Accuracy        float64   `json:"accuracy"`
	CreatedAt       time.Time `json:"created_at"`
}
