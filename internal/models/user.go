package models

import "errors"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password,omitempty" db:"password_hash" binding:"required"`
}

type UpdateNameInput struct {
	Name *string `json:"name"`
}

func (i UpdateNameInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
