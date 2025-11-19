package domain

import "errors"

type Todo struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
	Style       string `json:"style" db:"style"`
}

type CreateTodoInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateTodoInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
	Style       *string `json:"style"`
}

func (i UpdateTodoInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
