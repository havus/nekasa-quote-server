#!/bin/bash

# Load environment variables from .env file
if [ -f .env ]; then
  export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi

# Check if required environment variables are set
if [ -z "$DB_USER" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_HOST" ] || [ -z "$DB_PORT" ] || [ -z "$DB_NAME" ]; then
  echo "Database environment variables (DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME) must be set"
  exit 1
fi

# Run the migration command
$GOPATH/bin/migrate -database "mysql://$DB_USER:$DB_PASSWORD@tcp($DB_HOST:$DB_PORT)/$DB_NAME" -path internal/infrastructure/database/migrations "$@"
