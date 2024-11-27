package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

func (u *User) TableName() string {
	return "users"
}

var AllowedUserSortFields = map[string]string{
	"username": "username",
	"email":    "email",
	"created":  "created_at",
}
