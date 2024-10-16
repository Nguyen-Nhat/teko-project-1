-- name: GetBorrowBookDetailByBorrowId :many
SELECT *
FROM "Borrow_Book_Detail"
WHERE borrow_book_id = $1;