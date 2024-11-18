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
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	if post.Title == "" || post.Content == "" {
		return fmt.Errorf("title and content are required")
	}
	return s.repo.CreatePost(post)
}
