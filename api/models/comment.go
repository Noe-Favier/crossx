package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string `json:"content"`
	Post    Post   `json:"post" gorm:"foreignKey:PostID"`
	User    User   `json:"user" gorm:"foreignKey:UserID"`

	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}
