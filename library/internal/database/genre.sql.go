// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: genre.sql

package database

import (
	"context"
)

const createGenre = `-- name: CreateGenre :one
INSERT INTO "Genre" (name)
VALUES ($1)
RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateGenre(ctx context.Context, name string) (Genre, error) {
	row := q.db.QueryRowContext(ctx, createGenre, name)
	var i Genre
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAllGenres = `-- name: GetAllGenres :many
SELECT g.id, g.name, g.created_at, g.updated_at
FROM "Genre" as g
`

func (q *Queries) GetAllGenres(ctx context.Context) ([]Genre, error) {
	rows, err := q.db.QueryContext(ctx, getAllGenres)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Genre
	for rows.Next() {
		var i Genre
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const getGenreByBookID = `-- name: GetGenreByBookID :many
SELECT g.id, g.name, g.created_at, g.updated_at
FROM "Genre" as g INNER JOIN "Book_Genre" as bg
ON g.id = bg.genre_id
WHERE bg.book_id = $1
`

func (q *Queries) GetGenreByBookID(ctx context.Context, bookID int32) ([]Genre, error) {
	rows, err := q.db.QueryContext(ctx, getGenreByBookID, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Genre
	for rows.Next() {
		var i Genre
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const getGenreByID = `-- name: GetGenreByID :one
SELECT g.id, g.name, g.created_at, g.updated_at
FROM "Genre" as g
WHERE id = $1
`

func (q *Queries) GetGenreByID(ctx context.Context, id int32) (Genre, error) {
	row := q.db.QueryRowContext(ctx, getGenreByID, id)
	var i Genre
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPageGenres = `-- name: GetPageGenres :many
SELECT g.id, g.name, g.created_at, g.updated_at
FROM "Genre" as g
LIMIT $1
OFFSET $1 * ($2::INTEGER - 1)
`

type GetPageGenresParams struct {
	Size int32
	Page int32
}

func (q *Queries) GetPageGenres(ctx context.Context, arg GetPageGenresParams) ([]Genre, error) {
	rows, err := q.db.QueryContext(ctx, getPageGenres, arg.Size, arg.Page)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Genre
	for rows.Next() {
		var i Genre
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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
