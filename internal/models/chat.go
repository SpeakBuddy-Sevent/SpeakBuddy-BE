package models

import "time"

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Participants    []string           `bson:"participants" json:"participants"`
	LastMessageText string             `bson:"lastMessageText,omitempty" json:"lastMessageText"`
	LastMessageTime time.Time          `bson:"lastMessageTime,omitempty" json:"lastMessageTime"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatID    primitive.ObjectID `bson:"chatId" json:"chat_id"`
	SenderID  string             `bson:"senderId" json:"sender_id"`
	Text      string             `bson:"text" json:"text"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}
