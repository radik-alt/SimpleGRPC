# Архитектура проекта

## Схема работы с моделями

```
gRPC Request (authGrpc.GetUserRequest)
         ↓
    [Handler] ← конвертирует gRPC модели ↔ domain модели
         ↓
    [Service] ← бизнес-логика, работает с domain моделями
         ↓
  [Repository] ← работает с БД, возвращает domain модели
         ↓
      Database
```

## Слои приложения

### 1. **Handler (gRPC/HTTP)** - `week_1/grpc/internal/handler/`
- Принимает gRPC запросы (`authGrpc.GetUserRequest`)
- Вызывает Service
- Конвертирует domain модели в gRPC модели (`authGrpc.User`)
- Возвращает gRPC ответы

### 2. **Service** - `week_1/postgres/internal/service/`
- Бизнес-логика
- Валидация
- Работает только с domain моделями (`model.Auth`)
- НЕ знает о gRPC

### 3. **Repository** - `week_1/postgres/internal/repository/`
- Работа с БД
- Возвращает domain модели (`model.Auth`)
- НЕ знает о gRPC
- НЕ содержит бизнес-логику

### 4. **Domain Model** - `week_1/postgres/internal/model/`
- Основные модели приложения
- Используются во всех слоях (кроме handler)

### 5. **gRPC Model** - `week_1/grpc/pkg/auth_v1/`
- Автогенерированные protobuf модели
- Используются только в handler слое

## Пример потока данных

### GetUser(id: 1)

```go
// 1. Handler получает gRPC запрос
func (h *Handler) GetUser(ctx context.Context, req *authGrpc.GetUserRequest) (*authGrpc.User, error) {
    // req.GetId() = 1
    
    // 2. Вызываем Service (получаем domain модель)
    auth, err := h.authService.Get(ctx, req.GetId())
    // auth = &model.Auth{ID: 1, Name: "John", Password: "hash", CreatedAt: time.Now()}
    
    // 3. Конвертируем domain → gRPC модель
    return &authGrpc.User{
        Id:        auth.ID,        // domain → gRPC
        Name:      auth.Name,
        Password:  auth.Password,
        CreatedAt: timestamppb.New(auth.CreatedAt), // time.Time → timestamppb
    }, nil
}

// Service (работает с domain моделями)
func (s *authService) Get(ctx context.Context, id int64) (*model.Auth, error) {
    // Вызываем репозиторий
    return s.authRepo.Get(ctx, id)
}

// Repository (возвращает domain модель)
func (r *repo) Get(ctx context.Context, id int64) (*model.Auth, error) {
    var auth model.Auth
    // SQL запрос...
    err = r.pool.QueryRow(ctx, query, args...).Scan(
        &auth.ID,
        &auth.Name,
        &auth.Password,
        &auth.CreatedAt,
    )
    return &auth, nil
}
```

## Почему НЕ использовать gRPC модели в репозитории?

❌ **Плохо:**
```go
// Repository возвращает gRPC модель
func (r *repo) Get(ctx context.Context, id int64) (*authGrpc.User, error) {
    // Проблемы:
    // 1. Repository зависит от gRPC (нарушение разделения слоев)
    // 2. Нельзя переиспользовать для HTTP/CLI
    // 3. Сложное тестирование
}
```

✅ **Хорошо:**
```go
// Repository возвращает domain модель
func (r *repo) Get(ctx context.Context, id int64) (*model.Auth, error) {
    // Преимущества:
    // 1. Независимость от транспортного слоя (gRPC/HTTP/CLI)
    // 2. Простое тестирование
    // 3. Чистая архитектура
}
```

## Конвертация типов

### time.Time → timestamppb.Timestamp
```go
timestamppb.New(auth.CreatedAt)
```

### timestamppb.Timestamp → time.Time
```go
grpcUser.CreatedAt.AsTime()
```

## Структура файлов

```
week_1/
├── grpc/
│   ├── pkg/auth_v1/          # gRPC модели (автогенерированные)
│   └── internal/
│       └── handler/auth/     # gRPC handlers (конвертация gRPC ↔ domain)
│
└── postgres/
    └── internal/
        ├── model/            # Domain модели
        ├── repository/       # Работа с БД (domain модели)
        ├── service/          # Бизнес-логика (domain модели)
        └── converter/        # Хелперы для конвертации (опционально)
```

## Checklist

- [ ] Repository работает с domain моделями (`model.Auth`)
- [ ] Service работает с domain моделями
- [ ] Handler конвертирует gRPC ↔ domain модели
- [ ] Domain модели не зависят от gRPC
- [ ] Repository не знает о gRPC
- [ ] Service не знает о gRPC
