package service

import (
	"news/internal/repository"
	"news/models"
)

type Service struct {
	repo        *repository.Repository
	PostService *PostService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo:        repo,
		PostService: NewPostService(repo.PostRepository)}
}

type Post interface {
	CreatePost(post *models.Post) error
}

type User interface {
	CreateUser(user *models.User) error
}
