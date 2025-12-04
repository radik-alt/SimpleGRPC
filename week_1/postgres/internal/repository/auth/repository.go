package auth

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"postgres/internal/repository"
)

const (
	tableName = "auth"

	idColumn        = "id"
	nameColumn      = "name"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
)

type repo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool pgxpool.Pool) repository.AuthRepository {
	return &repo{pool: &pool}
}

func (r *repo) Create(ctx context.Context, username string, password string) (int64, error) {
	return 0, nil
}

func (r *repo) Get(ctx context.Context, id int64) (string, error) {
	return "", nil
}
