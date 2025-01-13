package repository

import (
	"fmt"
	"time"

	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (c *CommentPostgres) Create(comment models.Comment) (int, error) {
	comment.CreatedAt = time.Now()
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_id, user_id, content, created_at) values ($1, $2, $3, $4) RETURNING id", commentsTable)
	row := c.db.QueryRow(query, comment.PostID, comment.UserID, comment.Content, comment.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
