package request

type CreateSessionRequest struct {
	UserID     uint   `json:"user_id"`
	TargetText string `json:"target_text"`
	AudioURL   string `json:"audio_url"`
}
