package models

import (
	"time"
)

type User struct {
	ID                int       `json:"id"`
	Bio               string    `json:"bio"`
	Email             string    `json:"email"`
	Username          string    `json:"username"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PasswordHash      string    `json:"password_hash"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
}
