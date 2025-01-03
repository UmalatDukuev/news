package service

import (
	"github.com/UmalatDukuev/news"
	"github.com/UmalatDukuev/news/internal/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) CreatePost(post news.Post) (int, error) {
	return s.repo.CreatePost(post)
}
