package repository

import (
	"context"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func newUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{db}
}

func (u *userRepo) Create(ctx context.Context, name, username, email, hashedPassword, role string) (id int, err error) {
	ctx, cancel := context.WithTimeout(ctx, QueryDuration)
	defer cancel()

	query := "INSERT INTO users (name, username, email, password, role) VALUES ($1,$2,$3,$4,$5) RETURNING id;"

	err = u.db.QueryRowContext(ctx, query, name, username, email, hashedPassword, role).Scan(&id)

	return
}

func (u *userRepo) GetByCreds(ctx context.Context, username, hashedPassword string) (id int, name, email, role, created_at, updated_at string, err error) {
	ctx, cancel := context.WithTimeout(ctx, QueryDuration)
	defer cancel()

	query := "SELECT id, name, email, role, created_at, updated_at FROM users WHERE username=$1 AND password=$2;"

	err = u.db.QueryRowContext(ctx, query, username, hashedPassword).Scan(&id, &name, &email, &role, &created_at, &updated_at)

	return
}
