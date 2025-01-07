package repository

import (
	"fmt"

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
	fmt.Println("REPO: comment created")
	return 143, nil
}
