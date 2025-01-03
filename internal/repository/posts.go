package repository

import (
	"fmt"

	"github.com/UmalatDukuev/news"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

func (r *PostPostgres) CreatePost(post news.Post) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", postsTable)

	row := r.db.QueryRow(query, post.Title, post.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
