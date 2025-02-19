package services

import (
	"context"
	"ecom-mono-backend/internals/app/models"
	"ecom-mono-backend/internals/app/repository"
	"errors"
)

type userService struct {
	repo         repository.IUserRepo
	hashFunc     func(value string) string
	getTokenFunc func(tokenObj *models.AuthToken) (string, error)
}

func newUserService(repo repository.IUserRepo, hashFunc func(value string) string, getTokenFunc func(tokenObj *models.AuthToken) (string, error)) (IUserService, error) {
	if repo == nil {
		return nil, errors.New("nil value received for IUserRepo")
	}
	if hashFunc == nil {
		return nil, errors.New("nil value received for hashFunc")
	}
	if getTokenFunc == nil {
		return nil, errors.New("nil value received for getTokenFunc")
	}
	return &userService{repo, hashFunc, getTokenFunc}, nil
}

func (s *userService) Login(ctx context.Context, username, password string) (token string, tokenObj *models.AuthToken, err error) {
	hashedPassword := s.hashFunc(password)
	id, name, email, role, created_at, updated_at, err := s.repo.GetByCreds(ctx, username, hashedPassword)
	if err != nil {
		return
	}

	tokenObj = &models.AuthToken{
		ID: id, Name: name, Username: username, Email: email, Role: role,
		CreatedAt: created_at, UpdatedAt: updated_at,
	}

	token, err = s.getTokenFunc(tokenObj)

	return
}

func (s *userService) Signup(ctx context.Context, name, username, email, password, role string) (id int, err error) {
	hashedPassword := s.hashFunc(password)
	return s.repo.Create(ctx, name, username, email, hashedPassword, role)
}
