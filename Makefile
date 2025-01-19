include .env
export

MIGRATION_PATH = internal/adapter/storage/postgres/migrations
MAIN_FILE = cmd/http/main.go
TEST_PATH = internal/core/service

.PHONE: build run run-air stop install migrate-create migrate-up migrate-down docs-swag docs-fmt compose-build compose-run compose-stop test help

build: install
	@echo "Building the application"
	go build -o bin/server ${MAIN_FILE}

run: build
	@echo "Running the application"
	docker-compose up -d db
	go run ${MAIN_FILE}

run-air: build
	@echo "Running the application"
	docker-compose up -d db
	air -c air.toml

stop:
	@echo "Stopping the application"
	docker-compose down

install:
	go mod download
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

migrate-create:
	migrate create -ext sql -dir ${MIGRATION_PATH} -seq $(name)

migrate-up:
	migrate -path ./${MIGRATION_PATH} -database ${DATABASE_URL} -verbose up

migrate-down:
	migrate -path ./${MIGRATION_PATH} -database ${DATABASE_URL} -verbose down

docs-swag:
	swag init -g ${MAIN_FILE} --parseInternal true

docs-fmt:
	swag fmt ./...

compose-build:
	docker compose build

compose-run: compose-build
	docker compose up -d --wait

compose-stop:
	docker compose down

test:
	go test -v ./${TEST_PATH} -cover -coverprofile=coverage.out

coverage:
	@echo "🟢 Running coverage..."
	go tool cover -html=coverage.out

help:
	@echo "build: Build the application"
	@echo "compose-build: Build the docker compose"
	@echo "coverage: Show the coverage"
	@echo "docs-swag: Generate the swagger documentation"
	@echo "docs-fmt: Format the swagger documentation"
	@echo "help: Show this help message"
	@echo "install: Install the dependencies"
	@echo "migrate-create [name]: Create a new migration"
	@echo "migrate-down: Rollback the migrations"
	@echo "migrate-up: Run the migrations"
	@echo "run: Run the application"
	@echo "run-air: Run the application with air"
	@echo "compose-run: Run the docker compose"
	@echo "stop: Stop the application"
	@echo "compose-stop: Stop the docker compose"
	@echo "test: Run the tests"
