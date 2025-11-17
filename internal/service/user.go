package service

import (
	"github.com/kastuell/gotodoapp/internal/models"
	"github.com/kastuell/gotodoapp/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetMe(id int) (models.User, error) {
	return s.repo.GetById(id)
}
