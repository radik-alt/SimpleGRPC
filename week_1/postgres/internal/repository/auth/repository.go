package auth

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"postgres/internal/model"
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

func (r *repo) Create(ctx context.Context, auth *model.Auth) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, passwordColumn).
		Values(auth.Name, auth.Password).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	var id int64
	err = r.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Auth, error) {
	builder := sq.Select(idColumn, nameColumn, passwordColumn, createdAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var auth model.Auth
	err = r.pool.QueryRow(ctx, query, args...).Scan(
		&auth.ID,
		&auth.Name,
		&auth.Password,
		&auth.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &auth, nil
}
