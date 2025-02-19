package services

import (
	"context"
	"ecom-mono-backend/internals/app/models"
	"ecom-mono-backend/internals/app/repository"
	"ecom-mono-backend/internals/app/utils"
	"errors"
)

type IUserService interface {
	Signup(ctx context.Context, name, username, email, password, role string) (int, error)
	Login(ctx context.Context, username, password string) (token string, userDetails *models.AuthToken, err error)
}

type Service struct {
	User   IUserService
	Crypto utils.ICrypto
}

func NewService(repo *repository.Repository, crypto utils.ICrypto) (*Service, error) {
	if repo == nil {
		return nil, errors.New("nil value provided for Repository")
	}
	if crypto == nil {
		return nil, errors.New("nil value provided for ICrypto")
	}

	userService, err := newUserService(repo.User, crypto.Hash, crypto.GetTokenStr)
	if err != nil {
		return nil, err
	}
	return &Service{User: userService}, err
}
