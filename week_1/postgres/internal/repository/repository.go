package repository

import (
	"context"
	"postgres/internal/model"
)

type AuthRepository interface {
	Create(ctx context.Context, auth *model.Auth) (int64, error)
	Get(ctx context.Context, id int64) (*model.Auth, error)
}

type ChatRepository interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64)
}

type NoteRepository interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64)
}
