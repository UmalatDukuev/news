package service

import (
	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/models"
)

type Authorization interface {
	CreateUser(models.User) (int, error)
	GetUser(username string) (models.User, error)
	SignIn(username, password string) (models.User, error)
}

type Post interface {
	Create(models.Post) (int, error)
	GetById(int) (models.Post, error)
	GetAll() ([]models.Post, error)
	Update(post models.Post) error
}

type Like interface {
	Create(models.Like) (int, error)
	GetAllPostLikes(postID int) ([]models.LikeOnPost, error)
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
