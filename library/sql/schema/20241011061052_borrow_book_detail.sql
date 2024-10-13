-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Borrow_Book_Detail" (
    "id" integer PRIMARY KEY,
    "borrow_book_id" integer NOT NULL,
    "book_id" integer NOT NULL,
    "quantity" integer NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Borrow_Book_Detail";
-- +goose StatementEnd
