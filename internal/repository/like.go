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
	like.UserID = 1
	like.PostID = 2
	like.CreatedAt = time.Now()
	var id int
	fmt.Println("REPO: like created")
	query := fmt.Sprintf("INSERT INTO %s (post_id, user_id, created_at) values ($1, $2, $3)", likesTable)
	row := c.db.QueryRow(query, like.PostID, like.UserID, like.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return 143, nil
}
