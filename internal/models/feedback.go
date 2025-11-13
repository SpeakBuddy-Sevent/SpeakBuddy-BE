package models

import "time"

type Feedback struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SessionID uint      `json:"session_id"`
	AIModel   string    `json:"ai_model"`
	Feedback  string    `gorm:"type:text" json:"feedback"`
	CreatedAt time.Time `json:"created_at"`
}
