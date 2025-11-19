package service

import (
	"strconv"
	"time"

	"github.com/kastuell/gotodoapp/internal/auth"
	"github.com/kastuell/gotodoapp/internal/domain"
	"github.com/kastuell/gotodoapp/internal/hash"
	"github.com/kastuell/gotodoapp/internal/repository"
)

type AuthService struct {
	repo            repository.User
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	hasher          hash.PasswordHasher
}

type NewAuthServiceDeps struct {
	tokenManager    auth.TokenManager
	repo            repository.User
	hasher          hash.PasswordHasher
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAuthService(deps NewAuthServiceDeps) *AuthService {
	return &AuthService{
		repo:            deps.repo,
		hasher:          deps.hasher,
		accessTokenTTL:  deps.accessTokenTTL,
		refreshTokenTTL: deps.refreshTokenTTL,
	}
}

func (s *AuthService) Register(input domain.CreateUserInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Password)

	if err != nil {
		return Tokens{}, err
	}

	user := domain.User{
		Name:     input.Name,
		Username: input.Username,
		Password: passwordHash,
	}

	id, err := s.repo.Create(user)
	if err != nil {
		return Tokens{}, err
	}

	return s.createTokens(strconv.Itoa(id))
}

func (s *AuthService) Login(input domain.GetIdByCreditsInput) (Tokens, error) {
	passwordHash, err := s.hasher.Hash(input.Username)

	if err != nil {
		return Tokens{}, err
	}

	id, err := s.repo.GetIdByCredits(input.Username, passwordHash)

	if err != nil {
		return Tokens{}, err
	}

	return s.createTokens(strconv.Itoa(id))
}

func (s *AuthService) UpdateTokens(refreshToken string) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	userId, err := s.tokenManager.Parse(refreshToken)

	if err != nil {
		return Tokens{}, err
	}

	res.AccessToken, err = s.tokenManager.NewJWT(userId, s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewJWT(userId, s.refreshTokenTTL)
	if err != nil {
		return res, err
	}

	return res, err
}

func (s *AuthService) createTokens(userId string) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(userId, s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewJWT(userId, s.refreshTokenTTL)
	if err != nil {
		return res, err
	}

	return res, err
}
