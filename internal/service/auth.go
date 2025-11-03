package service

import (
	"github.com/UmalatDukuev/news/internal/errs"
	"github.com/UmalatDukuev/news/internal/repository"
	"github.com/UmalatDukuev/news/internal/utils"
	"github.com/UmalatDukuev/news/models"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	pass, err := utils.HashPassword(user.Password)
	if err != nil {
		return -1, errs.ErrHashingPass
	}
	user.Password = pass
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(username string) (models.User, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *AuthService) SignIn(username, password string) (models.User, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return models.User{}, err
	}

	ok, err := s.repo.CheckPassword(username, password)
	if err != nil {
		return models.User{}, err
	}
	if !ok {
		return models.User{}, errs.ErrInvalidCredentials
	}

	return user, nil
}
