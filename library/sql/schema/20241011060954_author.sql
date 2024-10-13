-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Author" (
    "id" integer PRIMARY KEY,
    "fullname" varchar NOT NULL DEFAULT '',
    "dob" date,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Author";
-- +goose StatementEnd
