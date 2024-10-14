-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Author" (
    "id" serial PRIMARY KEY,
    "fullname" varchar NOT NULL DEFAULT '',
    "dob" date,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Author";
-- +goose StatementEnd
