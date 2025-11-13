package request

type FeedbackRequest struct {
	SessionID uint   `json:"session_id" validate:"required"`
	InputText string `json:"input_text" validate:"required"`
	AIModel   string `json:"ai_model,omitempty"`
}
