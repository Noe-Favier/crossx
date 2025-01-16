package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`

	Content string `json:"content"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`

	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}
