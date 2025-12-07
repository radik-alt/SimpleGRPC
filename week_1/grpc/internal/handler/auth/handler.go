package auth

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	authGrpc "week_1/grpc/pkg/auth_v1"
	// Раскомментируйте когда подключите модуль postgres:
	// "postgres/internal/model"
	// "postgres/internal/service"
)

// AuthService интерфейс для работы с бизнес-логикой
// Когда подключите модуль postgres, замените на service.AuthService
type AuthService interface {
	// Create(ctx context.Context, auth *model.Auth) (int64, error)
	// Get(ctx context.Context, id int64) (*model.Auth, error)
}

// Handler обрабатывает gRPC запросы для auth
type Handler struct {
	authGrpc.UnimplementedAuthV1Server
	authService AuthService
}

// NewHandler создает новый handler
func NewHandler(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

// GetUser получает пользователя по ID
// Пример полной интеграции:
func (h *Handler) GetUser(ctx context.Context, req *authGrpc.GetUserRequest) (*authGrpc.User, error) {
	// 1. Вызываем service (получаем domain модель)
	// auth, err := h.authService.Get(ctx, req.GetId())
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get user: %w", err)
	// }

	// 2. Конвертируем domain модель в gRPC модель
	// return &authGrpc.User{
	// 	Id:        auth.ID,
	// 	Name:      auth.Name,
	// 	Password:  auth.Password,
	// 	CreatedAt: timestamppb.New(auth.CreatedAt),
	// }, nil

	// Временная заглушка
	return &authGrpc.User{
		Id:        req.GetId(),
		Name:      "example",
		Password:  "password",
		CreatedAt: timestamppb.Now(),
	}, nil
}

// CreateUser создает нового пользователя
func (h *Handler) CreateUser(ctx context.Context, req *authGrpc.CreateUserRequest) (*authGrpc.User, error) {
	// 1. Конвертируем gRPC запрос в domain модель (без ID и CreatedAt - они создаются в БД)
	// auth := &model.Auth{
	// 	Name:     req.GetName(),
	// 	Password: req.GetPassword(),
	// }

	// 2. Вызываем service
	// id, err := h.authService.Create(ctx, auth)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create user: %w", err)
	// }

	// 3. Получаем созданного пользователя
	// createdAuth, err := h.authService.Get(ctx, id)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get created user: %w", err)
	// }

	// 4. Конвертируем domain модель в gRPC
	// return &authGrpc.User{
	// 	Id:        createdAuth.ID,
	// 	Name:      createdAuth.Name,
	// 	Password:  createdAuth.Password,
	// 	CreatedAt: timestamppb.New(createdAuth.CreatedAt),
	// }, nil

	return nil, fmt.Errorf("not implemented")
}

// ListUsers возвращает список пользователей
func (h *Handler) ListUsers(ctx context.Context, req *authGrpc.ListUsersRequest) (*authGrpc.ListUsersResponse, error) {
	// Пример как вернуть список пользователей
	// users, err := h.authService.List(ctx, req.GetLimit(), req.GetOffset())
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to list users: %w", err)
	// }

	// Конвертируем каждого пользователя
	// grpcUsers := make([]*authGrpc.User, 0, len(users))
	// for _, user := range users {
	// 	grpcUsers = append(grpcUsers, &authGrpc.User{
	// 		Id:        user.ID,
	// 		Name:      user.Name,
	// 		Password:  user.Password,
	// 		CreatedAt: timestamppb.New(user.CreatedAt),
	// 	})
	// }

	// return &authGrpc.ListUsersResponse{
	// 	Users:      grpcUsers,
	// 	TotalCount: int64(len(users)),
	// }, nil

	return nil, fmt.Errorf("not implemented")
}

// UpdateUser обновляет пользователя
func (h *Handler) UpdateUser(ctx context.Context, req *authGrpc.UpdateUserRequest) (*authGrpc.User, error) {
	return nil, fmt.Errorf("not implemented")
}

// DeleteUser удаляет пользователя
func (h *Handler) DeleteUser(ctx context.Context, req *authGrpc.DeleteUserRequest) (*emptypb.Empty, error) {
	// Example:
	// err := h.authService.Delete(ctx, req.GetId())
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to delete user: %w", err)
	// }
	// return &emptypb.Empty{}, nil

	return nil, fmt.Errorf("not implemented")
}
