/*
POSTS
INT PRIMARY KEY AUTO INCREMENT,
INT NOT NULL,
Content TEXT NOT NULL,
VARCHAR(255),
CreatedAt TIMESTAMP DEFAULT CURRENT_TlMESTAtv1e
UpdatedAtTlMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE
CURRENT TIMESTAMP
FOREIGN KEY (Userld) REFERENCES Users(Userld) ON DELETE
CASCADE
*/

package models

import (
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	User      User      `json:"user"`
	Content   string    `json:"content"`
	MediaUrl  string    `json:"media_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID int `json:"-"`
}
