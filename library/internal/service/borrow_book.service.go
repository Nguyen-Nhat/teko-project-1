package service

import (
	"context"
	"database/sql"
	"fmt"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/internal/dto/res"
	"library/internal/util"
	"library/pkg/response"
)

type IBorrowBookService interface {
	GetBorrowBookDetails(ctx context.Context, data *req.BorrowBookDetailPageDto) (*res.PageResult, int, error)
	CreateBorrowBook(ctx context.Context, data *req.BorrowBookPostDto) (*res.BorrowBookDetailDto, int, error)
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

func (bbs *borrowBookService) GetBorrowBookDetails(ctx context.Context, data *req.BorrowBookDetailPageDto) (*res.PageResult, int, error) {
	params := database.GetBorrowBookDetailsParams{
		Page:      data.Page,
		Size:      data.Size,
		StudentID: util.ToNullString(data.StudentID),
		DayRange:  util.ToNullInt32(data.DayRange),
	}
	borrowRows, err := bbs.repository.GetBorrowBookDetails(ctx, params)
	if err != nil {
		return nil, response.CodeInternalServerError, err
	}
	totalPage := 0
	if len(borrowRows) > 0 {
		totalPage = int(borrowRows[0].TotalPage)
	}
	resultData := make([]res.BorrowItemWithBookInfoDto, len(borrowRows))
	for i, row := range borrowRows {
		resultData[i] = res.BorrowItemWithBookInfoDto{
			BorrowBookItemDto: res.BorrowBookItemDto{
				Id:           row.ID,
				BookID:       row.BookID,
				Quantity:     row.Quantity,
				BorrowBookID: row.BorrowBookID,
			},
			Title:      row.Title,
			StudentId:  row.StudentID,
			BorrowDate: row.BorrowDate,
			DueDate:    row.DueDate,
		}
	}
	result := res.PageResult{
		TotalPage: totalPage,
		Page:      int(data.Page),
		Size:      len(resultData),
		List:      resultData,
	}

	return &result, response.CodeSuccess, nil
}

func (bbs *borrowBookService) CreateBorrowBook(ctx context.Context, data *req.BorrowBookPostDto) (*res.BorrowBookDetailDto, int, error) {
	return withTransaction(ctx, bbs.db, bbs.repository, func(q *database.Queries) (*res.BorrowBookDetailDto, int, error) {
		params := database.CreateBorrowBookParams{
			StudentID: data.StudentId,
			DueDate:   data.DueDate,
		}
		borrowBook, err := q.CreateBorrowBook(ctx, params)
		if err != nil {
			return nil, response.CodeCannotCreateBorrowBook, err
		}

		bookIds := make([]int32, len(data.BorrowBookDetails))
		quantities := make([]int32, len(data.BorrowBookDetails))

		for i, borrowBookDetail := range data.BorrowBookDetails {
			bookIds[i] = borrowBookDetail.BookId
			quantities[i] = borrowBookDetail.Quantity
			params := database.UpdateBookQuantityWithBorrowBookParams{
				ID:       borrowBookDetail.BookId,
				Quantity: borrowBookDetail.Quantity,
			}
			rowAffected, err := q.UpdateBookQuantityWithBorrowBook(ctx, params)
			if err != nil {
				return nil, response.CodeInternalServerError, err
			}
			if rowAffected == 0 {
				return nil, response.CodeCannotCreateBorrowBook, fmt.Errorf("")
			}
		}

		detail, err := q.InsertBorrowBookDetails(ctx, database.InsertBorrowBookDetailsParams{
			BorrowBookID: borrowBook.ID,
			BookIds:      bookIds,
			Quantities:   quantities,
		})

		if err != nil {
			return nil, response.CodeCannotCreateBorrowBook, err
		}

		result := res.BorrowBookDetailDto{}
		result.FromModel(borrowBook, detail)
		return &result, response.CodeSuccess, nil
	})
}
func (bbs *borrowBookService) ReturnBorrowBook(ctx context.Context, id int) (interface{}, int, error) {
	return withTransaction(ctx, bbs.db, bbs.repository, func(q *database.Queries) (interface{}, int, error) {
		rowAffected, err := q.ReturnBorrowBook(ctx, int32(id))
		if err != nil || rowAffected == 0 {
			return nil, response.CodeCannotReturnBorrowBook, fmt.Errorf("")
		}
		borrowBookDetails, err := q.GetBorrowBookDetailByBorrowId(ctx, int32(id))

		if err != nil || borrowBookDetails == nil {
			return nil, response.CodeCannotReturnBorrowBook, fmt.Errorf("")
		}

		for _, borrowBookDetail := range borrowBookDetails {
			params := database.UpdateBookQuantityWithReturnBookParams{
				ID:       borrowBookDetail.BookID,
				Quantity: borrowBookDetail.Quantity,
			}
			rowAffected, err := q.UpdateBookQuantityWithReturnBook(ctx, params)
			if err != nil {
				return nil, response.CodeInternalServerError, err
			}
			if rowAffected == 0 {
				return nil, response.CodeCannotReturnBorrowBook, fmt.Errorf("")
			}

		}
		return nil, response.CodeSuccess, nil
	})
}
