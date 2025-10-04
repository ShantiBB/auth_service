run:
	go run ./cmd/auth/main.go

postgres-run:
	docker compose --env-file .env -f postgres-compose.yaml up -d
postgres-stop:
	docker compose --env-file .env -f postgres-compose.yaml down
