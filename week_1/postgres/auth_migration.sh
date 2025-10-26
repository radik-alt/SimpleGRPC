#!/bin/bash
source .env

export MIGRATION_DSN="host=auth-db port=5432 dbname=$POSTGRES_DB_AUTH user=$POSTGRES_USER_AUTH password=$POSTGRES_PASSWORD_AUTH sslmode=disable"

sleep 2 && goose -dir "${PG_DSN_AUTH}" postgres "${MIGRATION_DSN_AUTH}" up -v