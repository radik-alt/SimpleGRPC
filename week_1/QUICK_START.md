# Quick Start: Где писать Handler-ы?

## TL;DR

**Handler-ы пишутся в:** `week_1/grpc/internal/handler/<service>/handler.go`

**Пример:** `week_1/grpc/internal/handler/auth/handler.go` ✅ (уже создан)

---

## Визуальная структура

```
📁 week_1/grpc/
│
├── 📁 cmd/grpc_server/
│   └── 📄 main.go ← ЗДЕСЬ подключаются все handler-ы
│
├── 📁 internal/
│   └── 📁 handler/ ← ЗДЕСЬ живут все handler-ы
│       ├── 📁 auth/
│       │   └── 📄 handler.go ✅ (создан)
│       ├── 📁 chat/
│       │   └── 📄 handler.go ⚠️ (нужно создать)
│       └── 📁 note/
│           └── 📄 handler.go ⚠️ (нужно создать)
│
└── 📁 pkg/ ← автогенерированные protobuf (НЕ трогать)
```

---

## Как это работает?

### 1️⃣ Создаете Handler

**Файл:** `week_1/grpc/internal/handler/auth/handler.go`

```go
package auth

type Handler struct {
    authService service.AuthService  // ← получает Service
}

func NewHandler(authService service.AuthService) *Handler {
    return &Handler{authService: authService}
}

func (h *Handler) GetUser(ctx, req) (*authGrpc.User, error) {
    auth, _ := h.authService.Get(ctx, req.GetId())  // ← вызывает Service
    return ConvertToGrpc(auth), nil                 // ← конвертирует
}
```

### 2️⃣ Регистрируете в main.go

**Файл:** `week_1/grpc/cmd/grpc_server/main.go`

```go
func main() {
    // Инициализируем слои
    authRepo := authRepository.NewRepository(*pool)     // ← Repository
    authSvc := authService.NewService(authRepo)         // ← Service
    authHdl := authHandler.NewHandler(authSvc)          // ← Handler
    
    // Регистрируем handler
    s := grpc.NewServer()
    authGrpc.RegisterAuthV1Server(s, authHdl)  // ← ЗДЕСЬ!
    s.Serve(lis)
}
```

---

## Поток данных

```
1. gRPC Request
        ↓
2. Handler (grpc/internal/handler/auth/)
   • Принимает: authGrpc.GetUserRequest
   • Вызывает: authService.Get()
   • Возвращает: authGrpc.User
        ↓
3. Service (postgres/internal/service/auth/)
   • Бизнес-логика
   • Возвращает: model.Auth
        ↓
4. Repository (postgres/internal/repository/auth/)
   • SQL запросы
   • Возвращает: model.Auth
        ↓
5. Database
```

---

## Важные правила

✅ **Handler:**
- Знает о gRPC и domain моделях
- Конвертирует между ними
- Находится в `grpc/internal/handler/`

✅ **Service:**
- НЕ знает о gRPC
- Работает только с domain моделями
- Находится в `postgres/internal/service/`

✅ **Repository:**
- НЕ знает о gRPC
- НЕ знает о Service
- Работает только с БД и domain моделями
- Находится в `postgres/internal/repository/`

---

## Что уже сделано?

✅ Handler создан: `week_1/grpc/internal/handler/auth/handler.go`  
✅ Примеры всех методов: GetUser, CreateUser, ListUsers, etc.  
✅ Пример интеграции: `week_1/grpc/cmd/grpc_server/main_with_full_integration.go.example`

## Что делать дальше?

1. **Подключить модуль postgres в grpc**
   ```bash
   cd week_1/grpc
   # Добавьте в go.mod:
   # replace postgres => ../postgres
   ```

2. **Раскомментировать код в handler.go**
   - Импорты `postgres/internal/service` и `postgres/internal/model`
   - Реализацию методов

3. **Обновить main.go**
   - Использовать пример из `main_with_full_integration.go.example`

---

## Документация

📚 **Подробные гайды:**
- `HANDLERS_GUIDE.md` - детальное руководство по handler-ам
- `ARCHITECTURE.md` - полная архитектура проекта
- `README_MODELS.md` - работа с моделями
