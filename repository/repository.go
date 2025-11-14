package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/kastuell/gotodoapp/types"
)

type User interface {
	Create(user todo.User) (todo.User, error)
	GetById(id int) (todo.User, error)
	Delete(id int) (bool, error)
	Update(user todo.User) (todo.User, error)
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (todo.TodoItem, error)
	GetAllByUserId(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type TodoList interface {
	Create(todo todo.TodoList) (todo.TodoList, error)
	GetAll(id int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(id int) (bool, error)
	Update(todo.TodoList) (todo.TodoList, error)
}

type Repository struct {
	User
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoItem: NewTodoItemPostgres(db),
	}
}
