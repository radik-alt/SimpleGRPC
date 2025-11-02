# Подключаем переменные окружения из .env файла
include .env

# Директория для локальных бинарников
LOCAL_BIN:=$(CURDIR)/bin

# --- Переменные для Auth сервиса ---
# Путь к папке с миграциями для Auth
LOCAL_MIGRATION_DIR_AUTH=$(MIGRATION_DIR_AUTH)
# Строка подключения к Auth DB для локальных миграций (подключаемся к localhost и внешнему порту)
LOCAL_MIGRATION_DSN_AUTH="host=localhost port=$(POSTGRES_PORT) dbname=$(AUTH_DB_NAME) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

# --- Переменные для Chat сервиса ---
# Путь к папке с миграциями для Chat
LOCAL_MIGRATION_DIR_CHAT=$(MIGRATION_DIR_CHAT_SERVER)
# Строка подключения к Chat DB для локальных миграций
LOCAL_MIGRATION_DSN_CHAT="host=localhost port=$(POSTGRES_PORT) dbname=$(CHAT_DB_NAME) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"


.PHONY: install-deps
install-deps:
	# Устанавливаем goose, используя прямой доступ к репозиториям в обход прокси
	@echo "Installing goose..."
	@GOPROXY=direct GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

# --- Команды для Auth DB ---
.PHONY: local-auth-migration-status local-auth-migration-up local-auth-migration-down local-auth-migration-create

local-auth-migration-status:
	@echo "[AUTH] Checking migration status..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_AUTH} postgres ${LOCAL_MIGRATION_DSN_AUTH} status -v

local-auth-migration-up:
	@echo "[AUTH] Applying migrations..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_AUTH} postgres ${LOCAL_MIGRATION_DSN_AUTH} up -v

local-auth-migration-down:
	@echo "[AUTH] Rolling back one migration..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_AUTH} postgres ${LOCAL_MIGRATION_DSN_AUTH} down -v

local-auth-migration-create:
	@echo "[AUTH] Creating migration file: $(name)..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_AUTH} create $(name) sql


# --- Команды для Chat DB ---
.PHONY: local-chat-migration-status local-chat-migration-up local-chat-migration-down local-chat-migration-create

local-chat-migration-status:
	@echo "[CHAT] Checking migration status..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_CHAT} postgres ${LOCAL_MIGRATION_DSN_CHAT} status -v

local-chat-migration-up:
	@echo "[CHAT] Applying migrations..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_CHAT} postgres ${LOCAL_MIGRATION_DSN_CHAT} up -v

local-chat-migration-down:
	@echo "[CHAT] Rolling back one migration..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_CHAT} postgres ${LOCAL_MIGRATION_DSN_CHAT} down -v

local-chat-migration-create:
	@echo "[CHAT] Creating migration file: $(name)..."
	@$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR_CHAT} create $(name) sql
