package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kastuell/gotodoapp/internal/models"
)

type User interface {
	Create(user models.User) (int, error)
	GetIdByCredits(username, password_hash string) (int, error)
	GetById(id int) (models.User, error)
}

type TodoItem interface {
	Create(listId int, item models.Todo) (models.Todo, error)
	GetAllByUserId(userId, listId int) ([]models.Todo, error)
	GetById(userId, itemId int) (models.Todo, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type TodoList interface {
	Create(todo models.TodoList) (models.TodoList, error)
	GetAll(id int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(id int) (bool, error)
	Update(models.TodoList) (models.TodoList, error)
}

type Repository struct {
	User
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoItem: NewTodoItemPostgres(db),
		User:     NewUserPostgres(db),
	}
}
