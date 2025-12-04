package service

import "context"

type AuthService interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64) (string, error)
}

type ChatService interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64)
}
