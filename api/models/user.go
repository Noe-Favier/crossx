package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`

	Bio               string `json:"bio"`
	Email             string `json:"email" gorm:"not null;uniqueIndex"`
	Username          string `json:"username" gorm:"not null;uniqueIndex"`
	PasswordHash      string `json:"password_hash" gorm:"not null"`
	ProfilePictureUrl string `json:"profile_picture_url"`
}
