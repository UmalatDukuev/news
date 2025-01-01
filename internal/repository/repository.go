package repository

import (
	"github.com/UmalatDukuev/news"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
}

type Authorization interface {
	CreateUser(news.User) (int, error)
	GetUser(username, password string) (news.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
