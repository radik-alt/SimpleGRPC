# Где и как писать Handler-ы

## Структура проекта

```
week_1/
├── grpc/                              # gRPC модуль
│   ├── cmd/grpc_server/
│   │   └── main.go                    # ← ЗДЕСЬ регистрируются handler-ы
│   ├── internal/
│   │   └── handler/                   # ← ЗДЕСЬ пишутся handler-ы
│   │       ├── auth/
│   │       │   └── handler.go         # ✅ Handler для auth
│   │       ├── chat/
│   │       │   └── handler.go         # TODO: Handler для chat
│   │       └── note/
│   │           └── handler.go         # TODO: Handler для note
│   └── pkg/
│       └── auth_v1/                   # Автогенерированные protobuf
│           ├── auth.pb.go
│           └── auth_grpc.pb.go
│
└── postgres/                          # Postgres модуль (БД, Service, Repository)
    └── internal/
        ├── model/                     # Domain модели
        ├── repository/                # Работа с БД
        └── service/                   # Бизнес-логика
```

## Handler-ы: Где их писать?

### Ответ: `week_1/grpc/internal/handler/<service_name>/handler.go`

### Пример: Auth Handler

**Файл:** `week_1/grpc/internal/handler/auth/handler.go`

```go
package auth

import (
    "context"
    "google.golang.org/protobuf/types/known/timestamppb"
    
    authGrpc "week_1/grpc/pkg/auth_v1"
    "postgres/internal/model"
    "postgres/internal/service"
)

// Handler обрабатывает gRPC запросы
type Handler struct {
    authGrpc.UnimplementedAuthV1Server
    authService service.AuthService  // Service из модуля postgres
}

// NewHandler конструктор
func NewHandler(authService service.AuthService) *Handler {
    return &Handler{
        authService: authService,
    }
}

// GetUser - пример метода
func (h *Handler) GetUser(ctx context.Context, req *authGrpc.GetUserRequest) (*authGrpc.User, error) {
    // 1. Получаем domain модель из service
    auth, err := h.authService.Get(ctx, req.GetId())
    if err != nil {
        return nil, err
    }
    
    // 2. Конвертируем domain → gRPC
    return &authGrpc.User{
        Id:        auth.ID,
        Name:      auth.Name,
        Password:  auth.Password,
        CreatedAt: timestamppb.New(auth.CreatedAt),
    }, nil
}
```

## Как подключить Handler к серверу?

### Файл: `week_1/grpc/cmd/grpc_server/main.go`

```go
package main

import (
    "github.com/jackc/pgx/v4/pgxpool"
    "google.golang.org/grpc"
    
    // Handler
    authHandler "week_1/grpc/internal/handler/auth"
    
    // gRPC protobuf
    authGrpc "week_1/grpc/pkg/auth_v1"
    
    // Service и Repository из модуля postgres
    authService "postgres/internal/service/auth"
    authRepo "postgres/internal/repository/auth"
)

func main() {
    ctx := context.Background()
    
    // 1. Подключаемся к БД
    pool, err := pgxpool.Connect(ctx, dbDSN)
    defer pool.Close()
    
    // 2. Инициализируем слои (снизу вверх: Repository → Service → Handler)
    authRepository := authRepo.NewRepository(*pool)
    authSvc := authService.NewService(authRepository)
    authHdl := authHandler.NewHandler(authSvc)  // ← Handler получает Service
    
    // 3. Создаем gRPC сервер
    s := grpc.NewServer()
    
    // 4. Регистрируем handler
    authGrpc.RegisterAuthV1Server(s, authHdl)  // ← Регистрируем handler
    
    // 5. Запускаем сервер
    s.Serve(lis)
}
```

## Полный стек вызовов

```
Client Request (gRPC)
       ↓
[Handler] ← week_1/grpc/internal/handler/auth/handler.go
   │       • Получает gRPC запрос (authGrpc.GetUserRequest)
   │       • Вызывает Service
   │       • Конвертирует domain → gRPC (authGrpc.User)
   ↓
[Service] ← week_1/postgres/internal/service/auth/
   │       • Бизнес-логика
   │       • Валидация
   │       • Работает с model.Auth
   ↓
[Repository] ← week_1/postgres/internal/repository/auth/
   │       • SQL запросы
   │       • Возвращает model.Auth
   ↓
[Database] ← PostgreSQL
```

## Что делает каждый слой?

| Слой | Расположение | Ответственность |
|------|-------------|----------------|
| **Handler** | `grpc/internal/handler/` | Конвертация gRPC ↔ domain, маршрутизация |
| **Service** | `postgres/internal/service/` | Бизнес-логика, валидация |
| **Repository** | `postgres/internal/repository/` | Работа с БД |
| **Model** | `postgres/internal/model/` | Domain модели |

## Как добавить новый handler?

### Шаг 1: Создайте файл handler
```bash
mkdir -p week_1/grpc/internal/handler/chat
touch week_1/grpc/internal/handler/chat/handler.go
```

### Шаг 2: Напишите handler
```go
package chat

import (
    chatGrpc "week_1/grpc/pkg/chat_v1"
    "postgres/internal/service"
)

type Handler struct {
    chatGrpc.UnimplementedChatV1Server
    chatService service.ChatService
}

func NewHandler(chatService service.ChatService) *Handler {
    return &Handler{chatService: chatService}
}

// Реализуйте методы из protobuf...
```

### Шаг 3: Зарегистрируйте в main.go
```go
// Инициализация
chatRepo := chatRepository.NewRepository(*pool)
chatSvc := chatService.NewService(chatRepo)
chatHdl := chatHandler.NewHandler(chatSvc)

// Регистрация
chatGrpc.RegisterChatV1Server(s, chatHdl)
```

## FAQ

### Q: Можно ли писать handler в main.go?
**A:** Можно, но не рекомендуется. Лучше выносить в отдельные пакеты для:
- Читаемости кода
- Тестируемости
- Переиспользования
- Разделения ответственности

### Q: Должен ли handler знать о БД?
**A:** НЕТ! Handler работает только с Service. Service знает о Repository. Repository работает с БД.

### Q: Где конвертировать gRPC ↔ domain модели?
**A:** Только в Handler! Service и Repository работают исключительно с domain моделями.

### Q: Как тестировать handler?
**A:** Создайте mock Service и тестируйте конвертацию:
```go
// handler_test.go
mockService := &MockAuthService{}
handler := NewHandler(mockService)
// Тестируйте handler.GetUser(), handler.CreateUser() и т.д.
```

## Примеры

Полные примеры:
- ✅ `week_1/grpc/internal/handler/auth/handler.go` - готовый handler
- 📄 `week_1/grpc/cmd/grpc_server/main_with_full_integration.go.example` - пример интеграции
- 📚 `week_1/ARCHITECTURE.md` - архитектура проекта
