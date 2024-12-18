include .env
export

.PHONE: run stop install migrate-up migrate-down docs docs-fmt compose-build run-compose stop-compose test help

build: install
	@echo "Building the application"
	go build -o bin/server cmd/http/main.go

run: build
	@echo "Running the application"
	docker-compose up -d db
	go run cmd/http/main.go

run-air: build
	@echo "Running the application"
	docker-compose up -d db
	air -c air.toml

stop:
	@echo "Stopping the application"
	docker-compose down

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
	swag init -g cmd/http/main.go --parseInternal true

docs-fmt:
	swag fmt ./...

compose-build:
	docker compose build

run-compose: compose-build
	docker compose up -d --wait

stop-compose:
	docker compose down

test:
	go test -v ./...

help:
	@echo "install: Install the dependencies"
	@echo "build: Build the application"
	@echo "run: Run the application"
	@echo "stop: Stop the application"
	@echo "docs: Generate the swagger documentation"
	@echo "docs-fmt: Format the swagger documentation"
	@echo "compose-build: Build the docker compose"
	@echo "run-compose: Run the docker compose"
	@echo "stop-compose: Stop the docker compose"
	@echo "test: Run the tests"
	@echo "migrate-up: Run the migrations"
	@echo "migrate-down: Rollback the migrations"
	@echo "help: Show this help message"
