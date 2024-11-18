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
	fmt.Println(77777777777)

	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	if post.Title == "" || post.Content == "" {
		return fmt.Errorf("title and content are required")
	}
	fmt.Println(444444444)
	fmt.Println(post)
	err := s.repo.CreatePost(post)
	if err != nil {
		fmt.Println("Error in CreatePost:", err)
	}
	return err
}
