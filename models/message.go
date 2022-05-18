package models

import "time"

type Message struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	ChatID    string    `json:"chat_id" bson:"chat_id,omitempty"`
	UserID    string    `json:"user_id" bson:"user_id"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type CreateMessageRequest struct {
	ChatID  string `json:"chat_id" bson:"chat_id"`
	UserID  string `json:"user_id,swa" bson:"user_id" swaggerignore:"true"`
	Content string `json:"content" bson:"content"`
}
