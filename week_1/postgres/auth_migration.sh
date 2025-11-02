#!/bin/bash
source .env

sleep 2 && goose -dir "${MIGRATION_DIR_AUTH}" postgres "${MIGRATION_DSN_AUTH}" up -v