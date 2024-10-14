-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Borrow_Book_Detail" (
    "id" serial PRIMARY KEY,
    "borrow_book_id" integer NOT NULL,
    "book_id" integer NOT NULL,
    "quantity" integer NOT NULL DEFAULT 1
);
ALTER TABLE "Borrow_Book_Detail"
ADD CONSTRAINT "UQ_BookId_BorrowBookId" UNIQUE(book_id, borrow_book_id);

ALTER TABLE "Borrow_Book_Detail"
ADD CONSTRAINT "FK_Borrow_Book_Detail_To_Book"
FOREIGN KEY (book_id)
REFERENCES "Book"(id);

ALTER TABLE "Borrow_Book_Detail"
ADD CONSTRAINT "FK_Borrow_Book_Detail_To_Borrow_Book"
FOREIGN KEY (borrow_book_id)
REFERENCES "Borrow_Book"(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "Borrow_Book_Detail"
DROP CONSTRAINT "FK_Borrow_Book_Detail_To_Book";

ALTER TABLE "Borrow_Book_Detail"
DROP CONSTRAINT "FK_Borrow_Book_Detail_To_Borrow_Book";

ALTER TABLE "Borrow_Book_Detail"
DROP CONSTRAINT "UQ_BookId_BorrowBookId";

DROP TABLE IF EXISTS "Borrow_Book_Detail";
-- +goose StatementEnd
