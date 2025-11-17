package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	UsersTable      = "users"
	TodoListsTable  = "todo_lists"
	UsersListsTable = "users_lists"
	TodoItemsTable  = "todo_items"
	ListsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	initDB(db)

	return db, nil
}

func initDB(db *sqlx.DB) {
	createUsersTable(db)
	createTodoListsTable(db)
	createUsersListsTable(db)
	createTodoItemsTable(db)
	createListsItemsTable(db)
}

func createUsersTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id serial       not null unique,
		name          varchar(255) not null,
    	username      varchar(255) not null unique,
    	password_hash varchar(255) not null
	)`

	_, err := db.Exec(query)
	return err
}

func createTodoListsTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todo_lists (
		id          serial       not null unique,
		title       varchar(255) not null,
		description varchar(255)
	);`

	_, err := db.Exec(query)
	return err
}

func createUsersListsTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users_lists (
		id      serial                                           not null unique,
		user_id int references users (id) on delete cascade      not null,
		list_id int references todo_lists (id) on delete cascade not null
	);`

	_, err := db.Exec(query)
	return err
}

func createTodoItemsTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todo_items (
		id          serial       not null unique,
		title       varchar(255) not null,
		description varchar(255),
		done        boolean      not null default false
	);`

	_, err := db.Exec(query)
	return err
}

func createListsItemsTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS lists_items (
		id      serial                                           not null unique,
		item_id int references todo_items (id) on delete cascade not null,
		list_id int references todo_lists (id) on delete cascade not null
	);`

	_, err := db.Exec(query)
	return err
}
