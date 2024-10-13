GOOSE_DBSTRING ?= "postgres://root:root@127.0.0.1:5432/library?sslmode=disable"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= postgres

run:
	go run ./cmd/server/

goose_create:
	goose -dir $(GOOSE_MIGRATION_DIR) create $(name) sql

goose_up:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up

goose_down:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

goose_reset:
	goose -dir $(GOOSE_MIGRATION_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset

sqlgen:
	sqlc generate
swag:
	swag init -g ./cmd/server/main.go -o ./docs
.PHONY: run goose_create goose_up goose_down goose_reset sqlgen swag