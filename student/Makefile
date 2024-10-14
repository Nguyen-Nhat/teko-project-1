APP_NAME := server

GOOSE_DBSTRING ?= "postgres://root:root@127.0.0.1:5433/student?sslmode=disable"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= postgres

run:
	go run ./cmd/$(APP_NAME)/

goose_create:
	goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql

goose_up:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

goose_down:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

goose_reset:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset
swag:
	swag init -g ./cmd/server/main.go -o ./docs
wire:
	wire ./internal/wire
.PHONY: run goose_create goose_up goose_down goose_reset swag wire