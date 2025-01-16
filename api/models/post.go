package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`

	Content  string `json:"content"`
	MediaUrl string `json:"media_url"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`

	UserID int `json:"user_id"`
}
