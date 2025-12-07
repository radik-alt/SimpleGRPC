package auth

import (
	"context"
	"fmt"
	"postgres/internal/model"
)

func (serv *authService) Get(ctx context.Context, id int64) (*model.Auth, error) {
	auth, err := serv.authRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return auth, nil
}
