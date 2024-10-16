-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Borrow_Book" (
    "id" serial PRIMARY KEY,
    "student_id" varchar NOT NULL,
    "borrow_date" date NOT NULL DEFAULT now(),
    "due_date" date NOT NULL,
    "return_date" date,
    "is_return" boolean NOT NULL DEFAULT false,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Borrow_Book";
-- +goose StatementEnd
