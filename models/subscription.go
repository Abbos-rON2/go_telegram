package models

import "time"

type Subscription struct {
	ID        int       `json:"-"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}
