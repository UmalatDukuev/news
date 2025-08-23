package service

import (
	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/models"
)

type LikeService struct {
	repo repository.Like
}

func NewLikeService(repo repository.Like) *LikeService {
	return &LikeService{
		repo: repo,
	}
}

func (s *LikeService) Create(like models.Like) (int, error) {
	return s.repo.Create(like)
}

func (s *LikeService) GetAllPostLikes(postID int) ([]models.LikeOnPost, error) {
	return s.repo.GetAllPostLikes(postID)
}
