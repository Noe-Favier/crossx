package models

import (
	"time"
)

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	Post      Post      `json:"post"`
	User      User      `json:"user"`

	UserID int `json:"-"`
	PostID int `json:"-"`
}
