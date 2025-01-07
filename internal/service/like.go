package service

import (
	"fmt"

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
	fmt.Println("Service: like created")
	return s.repo.Create(like)

}
