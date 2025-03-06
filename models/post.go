package models

import "time"

type Post struct {
	ID          int       `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description" `
	AuthorID    int       `json:"author_id" db:"author_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
