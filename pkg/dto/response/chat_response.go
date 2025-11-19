package response

type ChatResponse struct {
	ID              string `json:"id"`
	Participants    []string `json:"participants"`
	LastMessageText string `json:"last_message"`
	LastMessageTime string `json:"last_time"`
	TherapistID  string    `json:"therapist_id"`
}

type MessageResponse struct {
	ID        string `json:"id"`
	ChatID    string `json:"chat_id"`
	SenderID  string `json:"sender_id"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
}
