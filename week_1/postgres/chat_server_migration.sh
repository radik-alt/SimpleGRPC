#!/bin/bash
# Загружаем переменные из .env файла
source .env

# Небольшая задержка на всякий случай, чтобы БД точно была готова
sleep 2

# Запускаем миграции
# -dir - указывает путь к папке с SQL файлами
# postgres - тип базы данных
# "${MIGRATION_DSN_CHAT_SERVER}" - строка подключения к БД
goose -dir "${MIGRATION_DIR_CHAT_SERVER}" postgres "${MIGRATION_DSN_CHAT_SERVER}" up -v
