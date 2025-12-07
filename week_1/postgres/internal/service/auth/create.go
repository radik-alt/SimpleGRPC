package auth

import (
	"context"
	"fmt"
	"postgres/internal/model"
)

func (serv *authService) Create(ctx context.Context, auth *model.Auth) (int64, error) {
	if auth.Name == "" {
		return 0, fmt.Errorf("username is required")
	}
	if auth.Password == "" {
		return 0, fmt.Errorf("password is required")
	}

	id, err := serv.authRepo.Create(ctx, auth)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}
