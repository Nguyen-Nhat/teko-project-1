-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Book" (
    "id" serial PRIMARY KEY,
    "title" varchar NOT NULL DEFAULT '',
    "published_year" integer NOT NULL,
    "available_quantity" integer NOT NULL DEFAULT 0,
    "borrow_quantity" integer NOT NULL DEFAULT 0,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Book";
-- +goose StatementEnd