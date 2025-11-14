package service

import (
	"github.com/kastuell/gotodoapp/repository"
	todo "github.com/kastuell/gotodoapp/types"
)

type Authorization interface {
	Register(user todo.User) (todo.User, error)
	Login()
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (todo.TodoItem, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Service struct {
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
