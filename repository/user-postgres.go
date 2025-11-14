package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/kastuell/gotodoapp/types"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user todo.User) (todo.User, error) {
	var newUser todo.User

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3)", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&newUser); err != nil {
		return newUser, err
	}

	return newUser, nil
}
