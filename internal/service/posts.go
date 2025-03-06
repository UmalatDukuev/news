package service

import (
	"fmt"

	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/models"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) Create(post models.Post) (int, error) {
	return s.repo.Create(post)
}

func (s *PostService) Update(post models.Post) error {

	existingPost, err := s.repo.GetById(post.ID)
	if err != nil {
		return err
	}

	if post.AuthorID != existingPost.AuthorID {
		return fmt.Errorf("you are not the author of this post")
	}

	if post.Title != "" {
		existingPost.Title = post.Title
	}
	if post.Description != "" {
		existingPost.Description = post.Description
	}

	return s.repo.Update(existingPost)
}

func (s *PostService) GetById(postID int) (models.Post, error) {
	return s.repo.GetById(postID)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}
