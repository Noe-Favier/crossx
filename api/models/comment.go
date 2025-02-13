package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content string    `json:"content"`
	Post    Post      `json:"post" gorm:"foreignKey:PostID"`
	User    User      `json:"user" gorm:"foreignKey:UserID"`
	Comment []Comment `json:"comments" gorm:"foreignKey:CommentID"`

	PostID uint `json:"post_id"`
	UserID uint `json:"user_id"`
}
