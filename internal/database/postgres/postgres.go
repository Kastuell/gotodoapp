package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	UsersTable      = "users"
	TodosTable      = "todos"
	TodosListsTable = "todos_lists"
	UsersListsTable = "users_lists"
	ListsTodosTable = "lists_todos"
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
	if err := createUserTable(db); err != nil {
		logrus.Error(err)
	}
	if err := createTodoListTable(db); err != nil {
		logrus.Error(err)
	}
	if err := createUsersListsTable(db); err != nil {
		logrus.Error(err)
	}
	if err := createTodoTable(db); err != nil {
		logrus.Error(err)
	}
	if err := createListsTodosTable(db); err != nil {
		logrus.Error(err)
	}
	if err := createSessionTable(db); err != nil {
		logrus.Error(err)
	}
}

func createUserTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id serial       not null unique,
		name          varchar(255) not null,
    	username      varchar(255) not null unique,
    	password_hash varchar(255) not null
	)`

	_, err := db.Exec(query)
	return err
}

func createTodoListTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos_lists (
		id          serial       not null unique,
		title       varchar(255) not null,
		description varchar(255)
	);`

	_, err := db.Exec(query)
	return err
}

func createUsersListsTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users_lists (
		id      serial                                          not null unique,
		user_id int references users (id) on delete cascade      not null,
		list_id int references todos_lists (id) on delete cascade not null
	);`

	_, err := db.Exec(query)
	return err
}

func createTodoTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id          serial       not null unique,
		title       varchar(255) not null,
		description varchar(255),
		done        boolean      not null default false,
		style       varchar(255) default 'default'
	);`

	_, err := db.Exec(query)
	return err
}

func createListsTodosTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS lists_todos (
		id      serial                                           not null unique,
		todo_id int references todos (id) on delete cascade not null,
		list_id int references todos_lists (id) on delete cascade not null
	);`

	_, err := db.Exec(query)
	return err
}

func createSessionTable(db *sqlx.DB) error {
	query := `CREATE TABLE IF NOT EXISTS lists_todos (
		id      serial                                           not null unique,
		user_id int not null,
		session_data json not null
	);`

	_, err := db.Exec(query)
	return err
}
