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
	Create(models.Post) (int, error)
	GetById(int) (models.Post, error)
	GetAll() ([]models.Post, error)
}

type Like interface {
	Create(models.Like) (int, error)
}

type Comment interface {
	Create(models.Comment) (int, error)
}

type Tag interface {
	Create(models.Tag) (int, error)
}

type Service struct {
	Authorization
	Post
	Like
	Comment
	Tag
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
		Like:          NewLikeService(repo.Like),
		Comment:       NewCommentService(repo.Comment),
		Tag:           NewTagService(repo.Tag),
	}
}
