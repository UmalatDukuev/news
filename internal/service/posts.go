package service

import (
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

func (s *PostService) GetById(postID int) (models.Post, error) {
	return s.repo.GetById(postID)
}

func (s *PostService) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}
