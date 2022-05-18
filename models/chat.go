package models

type Chat struct {
	ID        string    `json:"id"`
	Members   []string  `json:"user_id"`
	Messages  []Message `json:"content"`
	CreatedAt string    `json:"created_at"`
}

type CreateChatRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}
