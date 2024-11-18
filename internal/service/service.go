package service

import "news/internal/repository"

type Service struct {
	repo        *repository.Repository
	PostService *PostService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}
