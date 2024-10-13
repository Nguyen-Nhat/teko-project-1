-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Genre" (
    "id" integer PRIMARY KEY,
    "name" varchar NOT NULL DEFAULT '',
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Genre";
-- +goose StatementEnd
