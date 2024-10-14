-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Genre" (
    "id" serial PRIMARY KEY,
    "name" varchar NOT NULL DEFAULT '',
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Genre";
-- +goose StatementEnd
