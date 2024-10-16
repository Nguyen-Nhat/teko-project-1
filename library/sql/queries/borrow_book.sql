-- name: GetBorrowBookByID :one
SELECT bb.*
FROM "Borrow_Book" as bb
WHERE bb.id = $1;

-- name: CreateBorrowBook :exec
INSERT INTO "Borrow_Book"(student_id, due_date)
VALUES ($1, $2)
RETURNING *;

-- name: ReturnBorrowBook :exec
UPDATE "Borrow_Book"
SET is_return = true, updated_at = now(), return_date = now()
WHERE id = $1;