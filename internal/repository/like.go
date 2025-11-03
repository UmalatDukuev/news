package repository

import (
	"fmt"
	"time"

	"github.com/UmalatDukuev/news/internal/errs"
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

func (c *LikePostgres) GetAllPostLikes(postID int) ([]models.LikeOnPost, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE post_id = %d", likesTable, postID)
	rows, err := c.db.Query(query, postID)
	if err != nil {
		return nil, errs.ErrSQLQuery
	}
	defer rows.Close()
	for rows.Next() {
		var (
			postID     int
			userID     int
			created_at string
		)
		if err := rows.Scan(&postID, &userID, &created_at); err != nil {
			return nil, errs.ErrScanningRows
		}
	}
	likes := make([]models.LikeOnPost, 5)
	return likes, nil
}
