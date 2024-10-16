// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: borrow_book_detail.sql

package database

import (
	"context"
)

const getBorrowBookDetailByBorrowId = `-- name: GetBorrowBookDetailByBorrowId :many
SELECT id, borrow_book_id, book_id, quantity
FROM "Borrow_Book_Detail"
WHERE borrow_book_id = $1
`

func (q *Queries) GetBorrowBookDetailByBorrowId(ctx context.Context, borrowBookID int32) ([]BorrowBookDetail, error) {
	rows, err := q.db.QueryContext(ctx, getBorrowBookDetailByBorrowId, borrowBookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BorrowBookDetail
	for rows.Next() {
		var i BorrowBookDetail
		if err := rows.Scan(
			&i.ID,
			&i.BorrowBookID,
			&i.BookID,
			&i.Quantity,
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