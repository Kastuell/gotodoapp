package service

import (
	"github.com/kastuell/gotodoapp/internal/models"
	"github.com/kastuell/gotodoapp/internal/repository"
)

type User interface {
	GetMe(id int) (models.User, error)
}

type Auth interface {
	Register(user models.User) (string, error)
	Login(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoItem interface {
	Create(userId, listId int, item models.Todo) (models.Todo, error)
	GetAll(userId, listId int) ([]models.Todo, error)
	GetById(userId, itemId int) (models.Todo, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Service struct {
	TodoItem
	Auth
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
		Auth:     NewAuthService(repos.User),
		User:     NewUserService(repos.User),
	}
}
