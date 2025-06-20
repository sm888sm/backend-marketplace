.PHONY: run migrate createdb dropdb resetdb force-clean test help

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=marketplace

# Server
SERVER_PORT=8080
JWT_SECRET=pelindo888

help:
	@echo "Available targets:"
	@echo "  run        - Jalankan aplikasi Go"
	@echo "  migrate    - Jalankan migration ke database"
	@echo "  createdb   - Membuat database PostgreSQL"
	@echo "  dropdb     - Menghapus database PostgreSQL"
	@echo "  resetdb    - Drop lalu create database PostgreSQL"
	@echo "  force-clean - Paksa reset versi migration"
	@echo "  test       - Jalankan unit test"

run:
	go run cmd/main/main.go

migrate:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up

createdb:
	PGPASSWORD=$(DB_PASSWORD) createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME) || true

dropdb:
	PGPASSWORD=$(DB_PASSWORD) dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) --if-exists $(DB_NAME)

resetdb: dropdb createdb

force-clean:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force 1

test:
	go test ./...
