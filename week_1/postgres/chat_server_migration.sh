#!/bin/bash
source .env

export MIGRATION_DSN="host=chat-server-db port=5432 dbname=$POSTGRES_DB_CHAT_SERVER user=$POSTGRES_USER_CHAT_SERVER password=$POSTGRES_PASSWORD_CHAT_SERVER sslmode=disable"

sleep 2 && goose -dir "${PG_DSN_CHAT_SERVER}" postgres "${MIGRATION_DSN_CHAT_SERVER}" up -v