-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Book_Author" (
    "id" integer PRIMARY KEY,
    "book_id" integer NOT NULL,
    "author_id" integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Book_Author";
-- +goose StatementEnd
