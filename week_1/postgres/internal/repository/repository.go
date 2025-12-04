package repository

import "context"

type AuthRepository interface {
	Create(ctx context.Context, username string, password string) (int64, error)
	Get(ctx context.Context, id int64) (string, error)
}

type ChatRepository interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64)
}

type NoteRepository interface {
	Create(ctx context.Context)
	Get(ctx context.Context, id int64)
}
