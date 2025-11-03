package repository

import (
	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"
)

type TagPostgres struct {
	db *sqlx.DB
}

func NewTagPostgres(db *sqlx.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

func (c *TagPostgres) Create(tag models.Tag) (int, error) {
	return 143, nil
}
