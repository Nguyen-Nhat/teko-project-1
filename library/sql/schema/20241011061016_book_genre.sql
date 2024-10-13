-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Book_Genre" (
    "id" integer PRIMARY KEY,
    "book_id" integer NOT NULL,
    "genre_id" integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "Book_Genre";
-- +goose StatementEnd
