package repository

import (
	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Post
}

type Post interface {
	Create(models.Post) (int, error)
	GetById(int) (models.Post, error)
	GetAll() ([]models.Post, error)
}

type Authorization interface {
	CreateUser(models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Post:          NewPostPostgres(db),
	}
}
