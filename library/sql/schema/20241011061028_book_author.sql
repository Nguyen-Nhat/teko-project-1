-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Book_Author" (
    "id" serial PRIMARY KEY,
    "book_id" integer NOT NULL,
    "author_id" integer NOT NULL
);
ALTER TABLE "Book_Author"
ADD CONSTRAINT "UQ_BookId_AuthorId" UNIQUE(book_id, author_id);

ALTER TABLE "Book_Author"
ADD CONSTRAINT "FK_Book_Author_To_Book"
FOREIGN KEY (book_id)
REFERENCES "Book"(id);

ALTER TABLE "Book_Author"
ADD CONSTRAINT "FK_Book_Author_To_Author"
FOREIGN KEY (author_id)
REFERENCES "Author"(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "Book_Author"
DROP CONSTRAINT "FK_Book_Author_To_Book";

ALTER TABLE "Book_Author"
DROP CONSTRAINT "FK_Book_Author_To_Author";

ALTER TABLE "Book_Author"
DROP CONSTRAINT "UQ_BookId_AuthorId";

DROP TABLE IF EXISTS "Book_Author";
-- +goose StatementEnd
