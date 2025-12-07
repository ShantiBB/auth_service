#!/bin/bash
USER=$(yq e '.postgres.user' $CONFIG_PATH)
PASSWORD=$(yq e '.postgres.password' $CONFIG_PATH)
HOST=$(yq e '.postgres.host' $CONFIG_PATH)
PORT=$(yq e '.postgres.port' $CONFIG_PATH)
DB=$(yq e '.postgres.db' $CONFIG_PATH)
SSLMODE=$(yq e '.postgres.sslmode' $CONFIG_PATH)

DB_URL="postgres://$USER:$PASSWORD@$HOST:$PORT/$DB?sslmode=$SSLMODE"

migrate -path ./services/auth/migrations -database "$DB_URL" $1 # up / down
