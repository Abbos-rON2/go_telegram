package models

import "time"

type Message struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateMessageRequest struct {
	ChatID  string `json:"chat_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
