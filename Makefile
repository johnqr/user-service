.PHONY: build run test migrate-up migrate-down

build:
	go build ./...

run:
	go run ./cmd/api

test:
	go test ./...

migrate-up:
	@echo "Apply SQL in migrations/*.up.sql to your DB (use your migration tool)"

migrate-down:
	@echo "Apply SQL in migrations/*.down.sql to rollback"
