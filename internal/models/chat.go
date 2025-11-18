package models

import "time"

type Chat struct {
	ID                 string    `bson:"_id,omitempty" json:"id"`
	Participants       []string  `bson:"participants" json:"participants"`
	LastMessageText    string    `bson:"lastMessageText" json:"lastMessageText"`
	LastMessageTime    time.Time `bson:"lastMessageTime" json:"lastMessageTime"`
}

type Message struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	ChatID    string    `bson:"chatId" json:"chat_id"`
	SenderID  string    `bson:"senderId" json:"sender_id"`
	Text      string    `bson:"text" json:"text"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
}
