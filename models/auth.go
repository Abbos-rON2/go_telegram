package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Token struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
