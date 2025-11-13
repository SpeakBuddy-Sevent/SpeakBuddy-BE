package response

import "time"

type FeedbackResponse struct {
	ID        uint      `json:"id"`
	SessionID uint      `json:"session_id"`
	AIModel   string    `json:"ai_model"`
	Feedback  string    `json:"feedback"`
	CreatedAt time.Time `json:"created_at"`
}
