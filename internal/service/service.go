package service

import (
	"github.com/UmalatDukuev/news"
	"github.com/UmalatDukuev/news/internal/repository"
)

type Authorization interface {
	CreateUser(news.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
