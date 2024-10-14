-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Book_Genre" (
    "id" serial PRIMARY KEY,
    "book_id" integer NOT NULL,
    "genre_id" integer NOT NULL
);
ALTER TABLE "Book_Genre"
ADD CONSTRAINT "UQ_BookId_GenreId" UNIQUE(book_id, genre_id);

ALTER TABLE "Book_Genre"
ADD CONSTRAINT "FK_Book_Genre_To_Book"
FOREIGN KEY (book_id)
REFERENCES "Book"(id);

ALTER TABLE "Book_Genre"
ADD CONSTRAINT "FK_Book_Genre_To_Genre"
FOREIGN KEY (genre_id)
REFERENCES "Genre"(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "Book_Genre"
DROP CONSTRAINT "FK_Book_Genre_To_Book";

ALTER TABLE "Book_Genre"
DROP CONSTRAINT "FK_Book_Genre_To_Genre";

ALTER TABLE "Book_Genre"
DROP CONSTRAINT "UQ_BookId_GenreId";
DROP TABLE IF EXISTS "Book_Genre";
-- +goose StatementEnd
