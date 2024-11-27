package model

import (
	"github.com/google/uuid"
	"github.com/umardev500/pos-api/pkg"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash,omitempty"`
	pkg.TimeModel
}

func (u *User) TableName() string {
	return "users"
}

var AllowedUserSortFields = map[string]string{
	"username": "username",
	"email":    "email",
	"created":  "created_at",
}
