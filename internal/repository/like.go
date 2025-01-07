package repository

import (
	"fmt"

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
	fmt.Println("REPO: like created")
	return 143, nil
}
