package response

import "time"

type ExerciseItemResponse struct {
	ID         uint                      `json:"id"`
	ItemNumber int                       `json:"item_number"`
	TargetText string                    `json:"target_text"`
	Attempts   []ExerciseAttemptResponse `json:"attempts,omitempty"`
	CreatedAt  time.Time                 `json:"created_at"`
}

type ExerciseAttemptResponse struct {
	ID              uint      `json:"id"`
	TranscribedText string    `json:"transcribed_text"`
	Feedback        string    `json:"feedback"`
	AIModel         string    `json:"ai_model"`
	Accuracy        float64   `json:"accuracy"`
	CreatedAt       time.Time `json:"created_at"`
}

type ReadingExerciseResponse struct {
	ID          uint                   `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Items       []ExerciseItemResponse `json:"items,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
