package service

import (
	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/models"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Post interface {
	CreatePost(models.Post) (int, error)
}

type Service struct {
	Authorization
	Post
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
	}
}
