package request

type FeedbackRequest struct {
	SessionID  uint   `json:"session_id" validate:"required"`
	TargetText string `json:"target_text" validate:"required"`
	InputText  string `json:"input_text" validate:"required"`
	AIModel    string `json:"ai_model,omitempty"`
}
