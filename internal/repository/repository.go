package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kastuell/gotodoapp/internal/domain"
)

type User interface {
	Create(user domain.User) (int, error)
	GetIdByCredits(username, password_hash string) (int, error)
	GetById(id int) (domain.User, error)
}

type Todo interface {
	Create(listId int, item domain.Todo) (domain.Todo, error)
	GetAllByUserId(userId, listId int) ([]domain.Todo, error)
	GetById(userId, itemId int) (domain.Todo, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input domain.UpdateTodoInput) error
}

type TodoList interface {
	Create(todo domain.TodoList) (domain.TodoList, error)
	GetAll(id int) ([]domain.TodoList, error)
	GetById(userId, listId int) (domain.TodoList, error)
	Delete(id int) (bool, error)
	Update(domain.TodoList) (domain.TodoList, error)
}

type Repositories struct {
	User
	Todo
	TodoList
}

func NewRepository(db *sqlx.DB) *Repositories {
	return &Repositories{
		Todo: NewTodoItemPostgres(db),
		User: NewUserPostgres(db),
	}
}
