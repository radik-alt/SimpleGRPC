# Handler Layer

## Что это?

Handler слой отвечает за:
- 📥 Прием gRPC запросов
- 🔄 Конвертацию между gRPC и domain моделями
- 📤 Вызов Service слоя
- ✅ Возврат gRPC ответов

## Структура

```
handler/
├── auth/
│   └── handler.go    ✅ Готов
├── chat/
│   └── handler.go    TODO
└── note/
    └── handler.go    TODO
```

## Шаблон handler

```go
package <service_name>

import (
    "context"
    "<service>Grpc" "week_1/grpc/pkg/<service>_v1"
    "postgres/internal/service"
)

type Handler struct {
    <service>Grpc.Unimplemented<Service>V1Server
    <service>Service service.<Service>Service
}

func NewHandler(<service>Service service.<Service>Service) *Handler {
    return &Handler{<service>Service: <service>Service}
}

// Реализуйте методы из protobuf...
```

## Пример: Auth Handler

См. `auth/handler.go` - полностью реализованный пример.

## Важно!

❌ Handler НЕ должен:
- Работать напрямую с БД
- Содержать бизнес-логику
- Знать о Repository

✅ Handler должен:
- Только конвертировать модели
- Вызывать методы Service
- Обрабатывать gRPC специфичные вещи
