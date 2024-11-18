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
	fmt.Println(666666666)
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	fmt.Println(55555555555555555)
	if r.db == nil {
		fmt.Println("Database connection is nil!")
		return fmt.Errorf("database not initialized")
	}
	return nil
}
