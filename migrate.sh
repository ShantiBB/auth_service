#!/bin/bash
SERVICE=""
ACTION=""
STEPS=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --service)
      SERVICE="$2"
      shift 2
      ;;
    --action)
      ACTION="$2"
      shift 2
      ;;
    *)
      if [[ "$1" =~ ^[0-9]+$ ]]; then
        STEPS="$1"
      fi
      shift
      ;;
  esac
done

if [ -z "$SERVICE" ] || [ -z "$ACTION" ]; then
  echo "Usage: $0 --service <service> --action <up|down|force|version> [steps]"
  exit 1
fi

CONFIG_PATH="./services/$SERVICE/config/local.yaml"
USER=$(yq e '.postgres.user' "$CONFIG_PATH")
PASSWORD=$(yq e '.postgres.password' "$CONFIG_PATH")
HOST=$(yq e '.postgres.host' "$CONFIG_PATH")
PORT=$(yq e '.postgres.port' "$CONFIG_PATH")
DB=$(yq e '.postgres.db' "$CONFIG_PATH")
SSLMODE=$(yq e '.postgres.sslmode' "$CONFIG_PATH")

DB_URL="postgres://$USER:$PASSWORD@$HOST:$PORT/$DB?sslmode=$SSLMODE"
MIGRATION_PATH="./services/$SERVICE/migrations"

if [ -n "$STEPS" ]; then
  migrate -path "$MIGRATION_PATH" -database "$DB_URL" "$ACTION" "$STEPS"
else
  migrate -path "$MIGRATION_PATH" -database "$DB_URL" "$ACTION"
fi