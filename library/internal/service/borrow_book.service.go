package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/pkg/response"
	"time"
)

type IBorrowBookService interface {
	CreateBorrowBook(ctx context.Context, data *req.BorrowBookPostDto) (interface{}, int, error)
	ReturnBorrowBook(ctx context.Context, id int) (interface{}, int, error)
}

type borrowBookService struct {
	db         *sql.DB
	repository *database.Queries
}

func NewBorrowBookService() IBorrowBookService {
	return &borrowBookService{
		db:         global.Db,
		repository: database.New(global.Db),
	}
}

func (bbs *borrowBookService) CreateBorrowBook(ctx context.Context, data *req.BorrowBookPostDto) (interface{}, int, error) {
	if data.ReturnDate.Before(time.Now()) {
		return nil, response.CodeInvalidReturnBookDate, nil
	}
	return nil, response.CodeSuccess, nil
}
func (bbs *borrowBookService) ReturnBorrowBook(ctx context.Context, id int) (interface{}, int, error) {
	tx, err := bbs.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, response.CodeInternalServerError, err
	}
	q := bbs.repository.WithTx(tx)

	borrowBookDetails, err := q.GetBorrowBookDetailByBorrowId(ctx, int32(id))
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeBorrowBookNotFound, err
	}

	for _, borrowBookDetail := range borrowBookDetails {
		params := database.UpdateBookQuantityWithReturnBookParams{
			ID:       borrowBookDetail.BookID,
			Quantity: borrowBookDetail.Quantity,
		}
		if err := q.UpdateBookQuantityWithReturnBook(ctx, params); err != nil {
			_ = tx.Rollback()
			return nil, response.CodeCannotReturnBorrowBook, err
		}
	}

	err = q.ReturnBorrowBook(ctx, int32(id))
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeCannotReturnBorrowBook, err
	}

	_ = tx.Commit()
	return nil, response.CodeSuccess, nil
}
