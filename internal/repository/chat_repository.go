package repository

import (
	"context"
	"time"

	"speakbuddy/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository interface {
	FindOrCreateChat(userID, therapistID string) (*models.Chat, error)
	InsertMessage(msg *models.Message) error
	GetMessages(chatID string) ([]models.Message, error)
	GetUserChats(userID string) ([]models.Chat, error)
}

type chatRepository struct {
	chatCol    *mongo.Collection
	messageCol *mongo.Collection
}

func NewChatRepository(db *mongo.Database) ChatRepository {
	return &chatRepository{
		chatCol:    db.Collection("chats"),
		messageCol: db.Collection("messages"),
	}
}

func (r *chatRepository) FindOrCreateChat(userID, therapistID string) (*models.Chat, error) {
	ctx := context.Background()

	filter := bson.M{
		"participants": bson.M{"$all": []string{userID, therapistID}},
	}

	var chat models.Chat
	err := r.chatCol.FindOne(ctx, filter).Decode(&chat)

	if err == mongo.ErrNoDocuments {
		chat = models.Chat{
			Participants:    []string{userID, therapistID},
			LastMessageText: "",
			LastMessageTime: time.Now(),
		}

		res, err := r.chatCol.InsertOne(ctx, chat)
		if err != nil {
			return nil, err
		}

		chat.ID = res.InsertedID.(string)
		return &chat, nil
	}

	return &chat, err
}

func (r *chatRepository) InsertMessage(msg *models.Message) error {
	_, err := r.messageCol.InsertOne(context.Background(), msg)
	return err
}

func (r *chatRepository) GetMessages(chatID string) ([]models.Message, error) {
	cursor, err := r.messageCol.Find(context.Background(), bson.M{"chatId": chatID})
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	err = cursor.All(context.Background(), &messages)
	return messages, err
}

func (r *chatRepository) GetUserChats(userID string) ([]models.Chat, error) {
	cursor, err := r.chatCol.Find(context.Background(), bson.M{
		"participants": userID,
	})
	if err != nil {
		return nil, err
	}

	var chats []models.Chat
	err = cursor.All(context.Background(), &chats)
	return chats, err
}
