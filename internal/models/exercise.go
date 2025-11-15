package models

import "time"

// ReadingExercise = 1 latihan membaca (berisi beberapa soal)
type ReadingExercise struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `json:"user_id"`
	Title     string         `json:"title"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Items     []ExerciseItem `gorm:"foreignKey:ExerciseID" json:"items,omitempty"`
}

// ExerciseItem = 1 soal dalam latihan (punya target_text)
type ExerciseItem struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	ExerciseID uint              `json:"exercise_id"`
	ItemNumber int               `json:"item_number"` // 1-5
	TargetText string            `json:"target_text"`
	CreatedAt  time.Time         `json:"created_at"`
	Attempts   []ExerciseAttempt `gorm:"foreignKey:ItemID" json:"attempts,omitempty"`
}

// ExerciseAttempt = 1 kali attempt user untuk 1 soal (transcribe + feedback)
type ExerciseAttempt struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	ItemID          uint      `json:"item_id"`
	TranscribedText string    `json:"transcribed_text"`
	AIFeedback      string    `gorm:"type:text" json:"feedback"`
	AIModel         string    `json:"ai_model"`
	Accuracy        float64   `json:"accuracy"`
	CreatedAt       time.Time `json:"created_at"`
}
