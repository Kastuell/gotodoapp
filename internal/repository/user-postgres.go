package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kastuell/gotodoapp/internal/database/postgres"
	"github.com/kastuell/gotodoapp/internal/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", postgres.UsersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *UserPostgres) GetIdByCredits(username, password_hash string) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", postgres.UsersTable)

	row := r.db.QueryRow(query, username, password_hash)

	if err := row.Scan(&id); err != nil {
		return id, err
	}

	return id, nil
}

func (r *UserPostgres) GetById(id int) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT username, name FROM %s WHERE id = $1", postgres.UsersTable)

	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.Username, &user.Name); err != nil {
		return user, err
	}

	return user, nil
}
