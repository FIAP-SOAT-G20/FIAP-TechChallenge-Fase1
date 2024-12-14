include .env
export

run:
	@echo "Running the application"
	go run cmd/server/main.go

install:
	go mod download
	go install github.com/go-task/task/v3/cmd/task@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

migrate-up:
	migrate -path ./internal/adapter/storage/postgres/migrations -database ${DATABASE_URL} -verbose up

migrate-down:
	migrate -path ./internal/adapter/storage/postgres/migrations -database ${DATABASE_URL} -verbose down

docs:
	swag init -g cmd/server/main.go --parseInternal true

docs-fmt:
	swag fmt ./...

help:
	@echo "run: Run the application"
	@echo "install: Install the dependencies"
	@echo "migrate-up: Run the migrations"
	@echo "migrate-down: Rollback the migrations"
	@echo "help: Show this help message"
