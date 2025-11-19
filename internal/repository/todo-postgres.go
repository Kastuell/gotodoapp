package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kastuell/gotodoapp/internal/database/postgres"
	"github.com/kastuell/gotodoapp/internal/domain"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Create(listId int, input domain.CreateTodoInput) (domain.Todo, error) {
	tx, err := r.db.Begin()
	var todo domain.Todo
	if err != nil {
		return todo, err
	}

	createTodoQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2)", postgres.TodosTable)

	row := tx.QueryRow(createTodoQuery, *input.Title, *input.Description)
	err = row.Scan(&todo)
	if err != nil {
		tx.Rollback()
		return todo, err
	}

	createListsTodosQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", postgres.ListsTodosTable)
	_, err = tx.Exec(createListsTodosQuery, listId, todo.ID)
	if err != nil {
		tx.Rollback()
		return todo, err
	}

	return todo, tx.Commit()
}

func (r *TodoPostgres) GetAllByUserId(userId, listId int) ([]domain.Todo, error) {
	var items []domain.Todo

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE li.list_id = $1 AND ul.user_id = $2`,
		postgres.TodosTable, postgres.ListsTodosTable, postgres.UsersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoPostgres) GetById(userId, todoId int) (domain.Todo, error) {
	var item domain.Todo

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on li.item_id = ti.id
	 INNER JOIN %s ul on ul.list_id == li.list_id WHERE ti.id = $1 AND ul.user_id = $2`, postgres.TodosTable, postgres.ListsTodosTable, postgres.UsersListsTable)

	if err := r.db.Get(&item, query, todoId, userId); err != nil {
		return item, nil
	}

	return item, nil
}

func (r *TodoPostgres) Delete(userId, todoId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		postgres.TodosTable, postgres.ListsTodosTable, postgres.UsersListsTable)
	_, err := r.db.Exec(query, userId, todoId)
	return err
}

func (r *TodoPostgres) Update(userId, todoId int, input domain.UpdateTodoInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}

	if input.Style != nil {
		setValues = append(setValues, fmt.Sprintf("style=$%d", argId))
		args = append(args, *input.Style)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		postgres.TodosTable, setQuery, postgres.ListsTodosTable, postgres.UsersListsTable, argId, argId+1)
	args = append(args, userId, todoId)

	_, err := r.db.Exec(query, args...)
	return err
}
