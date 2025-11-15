package response

import "time"

type RecordingResponse struct {
	ID              uint      `json:"id"`
	SessionID       uint      `json:"session_id"`
	TranscribedText string    `json:"transcribed_text"`
	AIFeedback      string    `json:"feedback"`
	AIModel         string    `json:"ai_model"`
	Accuracy        float64   `json:"accuracy"`
	CreatedAt       time.Time `json:"created_at"`
}

type TranscribeAndAnalyzeResponse struct {
	TranscribedText string  `json:"transcribed_text"`
	TargetText      string  `json:"target_text"`
	RecordingID     uint    `json:"recording_id"`
	Feedback        string  `json:"feedback"`
	AIModel         string  `json:"ai_model"`
	Accuracy        float64 `json:"accuracy"`
}
