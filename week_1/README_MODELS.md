# Как добавить модели из gRPC в репозиторий

## Короткий ответ
**НЕ добавляйте gRPC модели в репозиторий!** Используйте domain модели.

## Правильная архитектура

### 1. Domain модель (используется везде, кроме handler)
```go
// week_1/postgres/internal/model/auth.go
type Auth struct {
    ID        int64     `db:"id"`
    Name      string    `db:"name"`
    Password  string    `db:"password"`
    CreatedAt time.Time `db:"created_at"`
}
```

### 2. Repository (работает с domain моделью)
```go
// week_1/postgres/internal/repository/auth/repository.go
type AuthRepository interface {
    Create(ctx context.Context, username string, password string) (int64, error)
    Get(ctx context.Context, id int64) (*model.Auth, error) // ← domain модель
}

func (r *repo) Get(ctx context.Context, id int64) (*model.Auth, error) {
    var auth model.Auth
    // SQL запрос...
    return &auth, nil
}
```

### 3. Service (работает с domain моделью)
```go
// week_1/postgres/internal/service/auth/get.go
func (s *authService) Get(ctx context.Context, id int64) (*model.Auth, error) {
    return s.authRepo.Get(ctx, id)
}
```

### 4. Handler (конвертирует domain → gRPC)
```go
// week_1/grpc/internal/handler/auth/handler.go
func (h *Handler) GetUser(ctx context.Context, req *authGrpc.GetUserRequest) (*authGrpc.User, error) {
    // 1. Получаем domain модель
    auth, err := h.authService.Get(ctx, req.GetId())
    
    // 2. Конвертируем domain → gRPC
    return &authGrpc.User{
        Id:        auth.ID,
        Name:      auth.Name,
        Password:  auth.Password,
        CreatedAt: timestamppb.New(auth.CreatedAt), // конвертация типа
    }, nil
}
```

## Конвертация типов

### Domain → gRPC
```go
grpcUser := &authGrpc.User{
    Id:        domainAuth.ID,
    Name:      domainAuth.Name,
    Password:  domainAuth.Password,
    CreatedAt: timestamppb.New(domainAuth.CreatedAt), // time.Time → Timestamp
}
```

### gRPC → Domain
```go
domainAuth := &model.Auth{
    ID:        grpcUser.Id,
    Name:      grpcUser.Name,
    Password:  grpcUser.Password,
    CreatedAt: grpcUser.CreatedAt.AsTime(), // Timestamp → time.Time
}
```

## Что было сделано

✅ Обновлен интерфейс `AuthRepository` для работы с `*model.Auth`
✅ Реализованы методы `Create` и `Get` в репозитории  
✅ Обновлен Service для работы с domain моделями  
✅ Создан пример Handler для gRPC  
✅ Создан конвертер (опционально)

## Проверка

Запустите, чтобы убедиться, что все компилируется:
```bash
cd week_1/postgres
go build ./...

cd ../grpc
go build ./...
```

## См. также
- `ARCHITECTURE.md` - подробная документация архитектуры
- `week_1/grpc/internal/handler/auth/handler.go` - полный пример handler
- `week_1/postgres/internal/repository/auth/repository.go` - реализация репозитория
