package repository

import (
	"fmt"
	"news/models"

	"github.com/jmoiron/sqlx"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	if db == nil {
		fmt.Println("Error: Database connection is nil in NewPostRepository!")
	}
	fmt.Println("PostRepository created successfully.")
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	if r.db == nil {
		return fmt.Errorf("database connection is nil")
	}

	query := `
		INSERT INTO posts (user_id, title, content, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id `

	fmt.Println("Executing query:", query)

	err := r.db.QueryRow(query, post.UserID, post.Title, post.Content, post.CreatedAt).Scan(&post.ID)
	if err != nil {
		fmt.Printf("Error inserting post: %v\n", err)
		return err
	}

	fmt.Println("Post created successfully:", post)
	return nil
}
