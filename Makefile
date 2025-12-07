mock-gen-auth:
	rm -rf services/auth/internal/mocks
	cd services/auth && mockery

swag-gen-auth:
	swag init -g services/auth/cmd/app/main.go --output services/auth/docs

gen-migrate: #name=<migration_name>
	migrate create -ext sql -dir ./services/auth/migrations -seq $(name)

add-migrate: #name=<migration_name>
	migrate create -ext sql -dir ./services/auth/migrations -seq $(name)

test-unit-auth-handler:
	go test ./services/auth/internal/http/handler/

run-%:
	CONFIG_PATH=$(shell pwd)/services/$*/config/local.yaml go run ./services/$*/cmd/app/main.go

build-%:
	go build -o build/$*-service services/$*/cmd/app/main.go

postgres-up:
	docker compose --env-file docker/.env -f docker/postgres.yaml up -d
postgres-stop:
	docker compose --env-file docker/.env -f docker/postgres.yaml down
