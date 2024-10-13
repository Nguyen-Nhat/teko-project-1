-- name: GetAllAuthors :many
SELECT id, fullname, dob, created_at, updated_at
FROM "Author";

-- name: GetAuthorByID :one
SELECT id, fullname, dob, created_at, updated_at
FROM "Author"
WHERE id = $1;
