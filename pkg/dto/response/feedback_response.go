package response

import "time"

type FeedbackResponse struct {
	ID              uint      `json:"id"`
	SessionID       uint      `json:"session_id"`
	TranscribedText string    `json:"transcribed_text"`
	TargetText      string    `json:"target_text"`
	AIModel         string    `json:"ai_model"`
	Feedback        string    `json:"feedback"`
	CreatedAt       time.Time `json:"created_at"`
}
