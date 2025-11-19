package services

import (
	"time"

	"speakbuddy/internal/models"
	"speakbuddy/internal/repository"
	"speakbuddy/pkg/dto/request"
)

type ChatService interface {
	SendMessage(userID, therapistID string, req request.SendMessageRequest) (*models.Message, error)
	GetChatMessages(chatID string) ([]models.Message, error)
	GetMyChats(userID string) ([]models.Chat, error)
}

type chatService struct {
	repo repository.ChatRepository
}

func NewChatService(repo repository.ChatRepository) ChatService {
	return &chatService{repo}
}

func (s *chatService) SendMessage(userID, therapistID string, req request.SendMessageRequest) (*models.Message, error) {
	chat, err := s.repo.FindOrCreateChat(userID, therapistID)
	if err != nil {
		return nil, err
	}

	message := &models.Message{
		ChatID:    chat.ID,
		SenderID:  userID,
		Text:      req.Text,
		Timestamp: time.Now().UTC(),
	}

	// insert message
	if err := s.repo.InsertMessage(message); err != nil {
		return nil, err
	}

	// update chat last message summary
	// chat.ID is primitive.ObjectID so convert to hex string for repo helper
	if err := s.repo.UpdateChatLastMessage(message.ChatID.Hex(), message.Text, message.Timestamp); err != nil {
		// log error but do not fail sending message; you can change behavior if desired
		// return nil, err
	}

	return message, nil
}

func (s *chatService) GetChatMessages(chatID string) ([]models.Message, error) {
	return s.repo.GetMessages(chatID)
}

func (s *chatService) GetMyChats(userID string) ([]models.Chat, error) {
	return s.repo.GetUserChats(userID)
}
