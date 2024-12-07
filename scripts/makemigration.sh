#!/bin/bash

# Check if migrate is installed
if ! command -v migrate &> /dev/null; then
  echo "Error: migrate CLI is not installed. Please install it first."
  exit 1
fi

# Function to print usage
usage() {
  echo "Usage: $0 <migration_name>"
  exit 1
}

# Check if the migration name is provided
if [ -z "$1" ]; then
  echo "Error: Migration name is required."
  usage
fi

MIGRATION_NAME="$1"

# Validate migration name (example: only allow alphanumeric and underscores)
if [[ ! "$MIGRATION_NAME" =~ ^[a-zA-Z0-9_]+$ ]]; then
  echo "Error: Migration name must contain only alphanumeric characters and underscores."
  usage
fi

# Define the directory where migrations are stored
MIGRATIONS_DIR="./migrations"

# Create the migrations directory if it doesn't exist
mkdir -p "$MIGRATIONS_DIR"

# Create the migration using the migrate CLI
migrate create -digits 3 -ext sql -dir "$MIGRATIONS_DIR" -seq "$MIGRATION_NAME"
