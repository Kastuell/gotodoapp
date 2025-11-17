package service

import (
	"github.com/kastuell/gotodoapp/internal/models"
	"github.com/kastuell/gotodoapp/internal/repository"
)

type TodoService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoService {
	return &TodoService{repo: repo, listRepo: listRepo}
}

func (s *TodoService) Create(userId, listId int, item models.Todo) (models.Todo, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return models.Todo{}, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoService) GetAll(userId, listId int) ([]models.Todo, error) {
	return s.repo.GetAllByUserId(userId, listId)
}

func (s *TodoService) GetById(userId, itemId int) (models.Todo, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoService) Update(userId, itemId int, input models.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
