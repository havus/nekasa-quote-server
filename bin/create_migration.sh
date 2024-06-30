#!/bin/bash

# Make sure this command is in the root of your project
# Use this command: ./bin/create_migration.sh <table_name>

if [ -z "$1" ]
then
  echo "Error: Migration name required"
  echo "Usage: ./bin/create_migration.sh <migration_name>"
  exit 1
fi

MIGRATION_NAME=$1
MIGRATION_PATH="internal/infrastructure/database/migrations"

$GOPATH/bin/migrate create -ext sql -dir $MIGRATION_PATH $MIGRATION_NAME
