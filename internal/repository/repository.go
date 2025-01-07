package repository

import (
	"github.com/UmalatDukuev/news/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Post
	Like
	Comment
	Tag
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

type Like interface {
	Create(like models.Like) (int, error)
}

type Comment interface {
	Create(comment models.Comment) (int, error)
}

type Tag interface {
	Create(comment models.Tag) (int, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Post:          NewPostPostgres(db),
		Like:          NewLikePostgres(db),
		Comment:       NewCommentPostgres(db),
		Tag:           NewTagPostgres(db),
	}
}
