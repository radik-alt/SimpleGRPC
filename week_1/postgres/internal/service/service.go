package service

import (
	"context"
	"postgres/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, auth *model.Auth) (int64, error)
	Get(ctx context.Context, id int64) (*model.Auth, error)
}

type ChatService interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64)
}
