package service

import (
	"fmt"

	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/models"
)

type TagService struct {
	repo repository.Tag
}

func NewTagService(repo repository.Tag) *TagService {
	return &TagService{
		repo: repo,
	}
}

func (s *TagService) Create(tag models.Tag) (int, error) {
	fmt.Println("service: tag created")
	return s.repo.Create(tag)
}
