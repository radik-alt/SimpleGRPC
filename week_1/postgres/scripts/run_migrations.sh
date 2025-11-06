#!/bin/bash
set -e

echo "Waiting for database $DB_HOST:$DB_PORT..."
sleep 3

echo "Running migrations from $MIGRATION_DIR..."
goose -dir "$MIGRATION_DIR" postgres "host=$DB_HOST port=$DB_PORT dbname=$DB_NAME user=$DB_USER password=$DB_PASSWORD sslmode=disable" up
