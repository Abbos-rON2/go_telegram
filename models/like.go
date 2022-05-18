package models

import "time"

type Like struct {
	ID        string    `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateLikeRequest struct {
	PostID int `json:"post_id"`
	UserID int `json:"-"`
}
