package request

type CreateSessionRequest struct {
	SessionID  uint   `form:"session_id" validate:"required"`
	TargetText string `form:"target_text" validate:"required"`
}
