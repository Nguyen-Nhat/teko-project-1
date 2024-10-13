-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Book" (
    "id" integer PRIMARY KEY,
    "title" varchar NOT NULL DEFAULT '',
    "published_year" integer,
    "available_quantity" integer NOT NULL DEFAULT 0,
    "borrow_quantity" integer NOT NULL DEFAULT 0,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Book";
-- +goose StatementEnd