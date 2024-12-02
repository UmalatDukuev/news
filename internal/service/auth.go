package service

import (
	"log/slog"
	"news/internal/repository"
	"news/models"
)

type UserService struct {
	repo *repository.PostRepository
}

func NewUserService(carRepo repository.) *CarService {
	log = log.With(slog.String("component", "car service"))

	return &CarService{
		log:     log,
		carRepo: carRepo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {

}
