package service

import (
	"fmt"
	"news/internal/repository"
	"news/models"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	if repo == nil {
		fmt.Println("Error: Repository is nil in NewPostService!")
	}
	fmt.Println("PostService created successfully.")
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	if post.Title == "" || post.Content == "" {
		return fmt.Errorf("title and content are required")
	}
	err := s.repo.CreatePost(post)
	if err != nil {
		fmt.Println("Error in CreatePost:", err)
	}
	return err
}
