package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Phone     string    `json:"phone" bson:"phone"`
	Bio       string    `json:"bio" bson:"bio"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type CreateUserRequest struct {
	Name     string `json:"name" bson:"name"`
	Phone    string `json:"phone" bson:"phone"`
	Password string `json:"password" bson:"password"`
}

type UserDTO struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Phone     string    `json:"phone" bson:"phone"`
	Bio       string    `json:"bio" bson:"bio"`
	Username  string    `json:"username" bson:"username"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
