package models

type Chat struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Members   []string  `json:"members" bson:"members"`
	Type      string    `json:"type" bson:"type"`
	Messages  []Message `json:"messages" bson:"messages"`
	CreatedAt string    `json:"created_at" bson:"created_at"`
}

type CreateChatRequest struct {
	UserID  string `json:"user_id" bson:"user_id"`
	Content string `json:"content" bson:"content"`
}
