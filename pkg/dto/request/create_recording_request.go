package request

type CreateRecordingRequest struct {
	SessionID uint   `json:"session_id" validate:"required"`
	AudioFile []byte `json:"audio_file"`
}
