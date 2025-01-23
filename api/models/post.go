package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Content  string `json:"content"`
	MediaUrl string `json:"media_url"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`

	UserID int `json:"user_id"`
}
