-- name: GetBorrowBookDetailByBorrowId :many
SELECT *
FROM "Borrow_Book_Detail"
WHERE borrow_book_id = $1;

-- name: InsertBorrowBookDetails :many
INSERT INTO "Borrow_Book_Detail"(borrow_book_id,book_id,quantity)
VALUES(sqlc.arg(borrow_book_id), unnest(sqlc.arg(book_ids)::int[]), unnest(sqlc.arg(quantities)::int[]))
RETURNING *;