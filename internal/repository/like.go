package repository

import (
	"fmt"
	"time"

	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"
)

type LikePostgres struct {
	db *sqlx.DB
}

func NewLikePostgres(db *sqlx.DB) *LikePostgres {
	return &LikePostgres{db: db}
}

func (c *LikePostgres) Create(like models.Like) (int, error) {
	like.CreatedAt = time.Now()
	var id int
	query := fmt.Sprintf("INSERT INTO %s (post_id, user_id, created_at) values ($1, $2, $3) RETURNING id", likesTable)
	row := c.db.QueryRow(query, like.PostID, like.UserID, like.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
