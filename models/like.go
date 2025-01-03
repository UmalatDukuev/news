package models

import "time"

type Like struct {
	ID        int       `json:"id" db:"id"`
	PostID    int       `json:"post_id" db:"post_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Type      string    `json:"type" db:"type"` // "like" or "dislike"
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
