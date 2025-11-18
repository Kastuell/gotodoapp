package service

import (
	"context"
	"time"

	"github.com/kastuell/gotodoapp/internal/auth"
	"github.com/kastuell/gotodoapp/internal/domain"
	"github.com/kastuell/gotodoapp/internal/hash"
	"github.com/kastuell/gotodoapp/internal/repository"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type User interface {
	GetMe(id int) (domain.User, error)
}

type Auth interface {
	Register(ctx context.Context, input domain.User) (Tokens, error)
	Login(ctx context.Context, username, password string) (Tokens, error)
}

type Todo interface {
	Create(userId, listId int, item domain.Todo) (domain.Todo, error)
	GetAll(userId, listId int) ([]domain.Todo, error)
	GetById(userId, itemId int) (domain.Todo, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input domain.UpdateTodoInput) error
}

type Services struct {
	Todo
	Auth
	User
}

type NewServiceDeps struct {
	Repos           *repository.Repositories
	TokenManager    auth.TokenManager
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	Hasher          hash.PasswordHasher
}

func NewService(deps NewServiceDeps) *Services {
	return &Services{
		Todo: NewTodoService(deps.Repos.Todo, deps.Repos.TodoList),
		Auth: NewAuthService(NewAuthServiceDeps{
			tokenManager:    deps.TokenManager,
			repo:            deps.Repos.User,
			hasher:          deps.Hasher,
			accessTokenTTL:  deps.AccessTokenTTL,
			refreshTokenTTL: deps.RefreshTokenTTL,
		}),
		User: NewUserService(deps.Repos.User),
	}
}
