-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Borrow_Book" (
    "id" integer PRIMARY KEY,
    "student_id" integer NOT NULL,
    "borrow_date" date NOT NULL,
    "return_date" date NOT NULL,
    "is_return" boolean NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Borrow_Book";
-- +goose StatementEnd