package auth

import (
	"postgres/internal/repository"
	"postgres/internal/service"
)

type authService struct {
	authRepo repository.AuthRepository
}

func NewService(
	authRepo repository.AuthRepository,
) service.AuthService {
	return &authService{authRepo: authRepo}
}
