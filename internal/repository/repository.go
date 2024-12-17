package repository

import (
	"fmt"
	"news/models"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	PostRepository *PostRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	if db == nil {
		fmt.Println("Error: DB is nil in NewRepository!")
	}
	fmt.Println("NewRepository called.")
	return &Repository{
		PostRepository: NewPostRepository(db),
	}
}

type User interface {
	CreateUser(user *models.User) error
}

type Post interface {
	CreatePost(post *models.Post) error
}
