package repository

import (
	"context"
	"time"

	"speakbuddy/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatRepository interface {
	FindOrCreateChat(userID, therapistID string) (*models.Chat, error)
	InsertMessage(msg *models.Message) error
	GetMessages(chatID string) ([]models.Message, error)
	GetUserChats(userID string) ([]models.Chat, error)
	UpdateChatLastMessage(chatID string, text string, ts time.Time) error
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
		"participants": bson.M{"$all": bson.A{userID, therapistID}},
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

		oid, ok := res.InsertedID.(primitive.ObjectID)
		if ok {
			chat.ID = oid
		}
		return &chat, nil
	}
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

func (r *chatRepository) InsertMessage(msg *models.Message) error {
	_, err := r.messageCol.InsertOne(context.Background(), msg)
	return err
}

func (r *chatRepository) GetMessages(chatID string) ([]models.Message, error) {
	ctx := context.Background()
	oid, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return nil, err
	}

	cursor, err := r.messageCol.Find(ctx, bson.M{"chatId": oid})
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	err = cursor.All(ctx, &messages)
	return messages, err
}

func (r *chatRepository) GetUserChats(userID string) ([]models.Chat, error) {
	ctx := context.Background()
	// Cari chat yang participants array mengandung userID
	cursor, err := r.chatCol.Find(ctx, bson.M{
		"participants": bson.M{"$in": bson.A{userID}},
	})
	if err != nil {
		return nil, err
	}

	var chats []models.Chat
	err = cursor.All(ctx, &chats)
	return chats, err
}

func (r *chatRepository) UpdateChatLastMessage(chatID string, text string, ts time.Time) error {
	ctx := context.Background()
	oid, err := primitive.ObjectIDFromHex(chatID)
	if err != nil {
		return err
	}

	_, err = r.chatCol.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": bson.M{
			"lastMessageText": text,
			"lastMessageTime": ts,
		}},
	)
	return err
}
