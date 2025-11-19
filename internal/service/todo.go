package service

import (
	"github.com/kastuell/gotodoapp/internal/domain"
	"github.com/kastuell/gotodoapp/internal/repository"
)

type TodoService struct {
	repo     repository.Todo
	listRepo repository.TodoList
}

func NewTodoService(repo repository.Todo, listRepo repository.TodoList) *TodoService {
	return &TodoService{repo: repo, listRepo: listRepo}
}

func (s *TodoService) Create(userId, listId int, input domain.CreateTodoInput) (domain.Todo, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return domain.Todo{}, err
	}

	return s.repo.Create(listId, input)
}

func (s *TodoService) GetAll(userId, listId int) ([]domain.Todo, error) {
	return s.repo.GetAllByUserId(userId, listId)
}

func (s *TodoService) GetById(userId, todoId int) (domain.Todo, error) {
	return s.repo.GetById(userId, todoId)
}

func (s *TodoService) Delete(userId, todoId int) error {
	return s.repo.Delete(userId, todoId)
}

func (s *TodoService) Update(userId, todoId int, input domain.UpdateTodoInput) error {
	return s.repo.Update(userId, todoId, input)
}
