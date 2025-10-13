FROM golang:1.24.1-alpine AS builder

WORKDIR /week_1/grpc/source
COPY . .

RUN echo "=== ДЕТАЛЬНАЯ ОТЛАДКА ==="
RUN pwd
RUN echo "--- Полное содержимое ---"
RUN ls -la

RUN echo "Начинаем загрузку модулей..."
RUN go mod download
RUN echo "Модули загружены успешно!"

RUN echo "Собираем приложение..."
RUN go build -o ./crud_server cmd/grpc_server/main.go
RUN echo "Сборка завершена!"

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /week_1/grpc/source/crud_server .

CMD ["./crud_server"]