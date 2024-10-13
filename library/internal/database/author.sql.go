// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: author.sql

package database

import (
	"context"
)

const getAllAuthors = `-- name: GetAllAuthors :many
SELECT id, fullname, dob, created_at, updated_at
FROM "Author"
`

func (q *Queries) GetAllAuthors(ctx context.Context) ([]Author, error) {
	rows, err := q.db.QueryContext(ctx, getAllAuthors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Author
	for rows.Next() {
		var i Author
		if err := rows.Scan(
			&i.ID,
			&i.Fullname,
			&i.Dob,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAuthorByID = `-- name: GetAuthorByID :one
SELECT id, fullname, dob, created_at, updated_at
FROM "Author"
WHERE id = $1
`

func (q *Queries) GetAuthorByID(ctx context.Context, id int32) (Author, error) {
	row := q.db.QueryRowContext(ctx, getAuthorByID, id)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Dob,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}