-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "universities" (
    "id" integer PRIMARY KEY,
    "name" varchar NOT NULL DEFAULT '',
    "establishment_year" integer NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "universities";
-- +goose StatementEnd
