package service

import (
	"fmt"

	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/models"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) Create(comment models.Comment) (int, error) {
	fmt.Println("Service: comment created")
	return s.repo.Create(comment)

}
