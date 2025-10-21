package models

import "time"

type Session struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	TargetText  string    `json:"target_text"`
	Transcript  string    `json:"transcript"`
	Accuracy    float64   `json:"accuracy"`
	AudioURL    string    `json:"audio_url"`
	CreatedAt   time.Time `json:"created_at"`
}
