package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Bio               string `json:"bio"`
	Email             string `json:"-" gorm:"not null;uniqueIndex"`
	Username          string `json:"username" gorm:"not null;uniqueIndex"`
	PasswordHash      string `json:"-" gorm:"not null"`
	ProfilePictureUrl string `json:"profile_picture_url"`

	Posts    []Post    `json:"-"`
	Comments []Comment `json:"-"`
	Likes    []Post    `json:"-"`
}
