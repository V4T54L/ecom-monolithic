package repository

import (
	"context"
	"database/sql"
	"time"
)

const (
	QueryDuration = time.Second * 5
)

type IUserRepo interface {
	Create(ctx context.Context, name, username, email, hashedPassword, role string) (int, error)
	GetByCreds(ctx context.Context, username, hashedPassword string) (id int, name, email, role, created_at, updated_at string, err error)
}

type Repository struct {
	User IUserRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: newUserRepo(db),
	}
}
