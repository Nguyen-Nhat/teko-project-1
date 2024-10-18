-- name: GetBorrowBookDetailByBorrowId :many
SELECT *
FROM "Borrow_Book_Detail"
WHERE borrow_book_id = $1;

-- name: InsertBorrowBookDetails :many
INSERT INTO "Borrow_Book_Detail"(borrow_book_id,book_id,quantity)
VALUES(sqlc.arg(borrow_book_id), unnest(sqlc.arg(book_ids)::int[]), unnest(sqlc.arg(quantities)::int[]))
RETURNING *;

-- name: GetBorrowBookDetails :many
SELECT bbd.*, b.title, bb.student_id, bb.borrow_date , bb.due_date, CEIL(COUNT(*) OVER() / (1.0 * sqlc.arg(size))) AS total_page
FROM "Borrow_Book" bb INNER JOIN "Borrow_Book_Detail" bbd ON bb.id = bbd.borrow_book_id
INNER JOIN "Book" b ON bbd.book_id = b.id
WHERE bb.is_return = false
AND bb.student_id = sqlc.narg(student_id)
AND (EXTRACT(DAY FROM AGE(bb.due_date, NOW())) >= 0
         AND EXTRACT(DAY FROM AGE(bb.due_date, NOW())) <= sqlc.narg(day_range)::int
         OR sqlc.narg(day_range)::int IS NULL)
LIMIT sqlc.arg(size)
OFFSET sqlc.arg(size) * (sqlc.arg(page)::int - 1);