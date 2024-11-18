package repository

import (
	"database/sql"
	"fmt"
	"news/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	fmt.Println(55555555555555555)
	return nil
}
