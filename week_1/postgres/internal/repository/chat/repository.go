package chat

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"postgres/internal/repository"
)

type repo struct {
	pool *pgxpool.Pool
}

func NewChatRepository(pgxpool *pgxpool.Pool) repository.ChatRepository {
	return &repo{pgxpool}
}

func (r *repo) Create(ctx context.Context) {}

func (r *repo) Get(ctx context.Context, id int64) {}
