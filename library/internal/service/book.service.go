package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/request"
	"library/pkg/response"
)

type IBookService interface {
	CreateBook(ctx context.Context, data *request.BookPostDto) (*database.Book, int, error)
	AddGenreToBook(ctx context.Context, bookId int, genreId int) (*database.BookGenre, int, error)
	AddAuthorToBook(ctx context.Context, bookId int, authorId int) (*database.BookAuthor, int, error)
}

type bookService struct {
	db         *sql.DB
	repository *database.Queries
}

func NewBookService() IBookService {
	return &bookService{
		db:         global.Db,
		repository: database.New(global.Db),
	}
}
func (bs *bookService) CreateBook(ctx context.Context, data *request.BookPostDto) (*database.Book, int, error) {
	tx, _ := bs.db.BeginTx(ctx, nil)
	q := bs.repository.WithTx(tx)
	params := database.CreateBookParams{
		Title:         data.Title,
		PublishedYear: data.PublishedYear,
	}
	result, err := q.CreateBook(ctx, params)
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeCannotCreateBook, err
	}
	_ = tx.Commit()
	return &result, response.CodeCreated, nil
}
func (bs *bookService) AddGenreToBook(ctx context.Context, bookId int, genreId int) (*database.BookGenre, int, error) {
	tx, _ := bs.db.BeginTx(ctx, nil)
	q := bs.repository.WithTx(tx)

	if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
		return nil, response.CodeBookNotFound, err
	}

	if _, err := q.GetGenreByID(ctx, int32(genreId)); err != nil {
		return nil, response.CodeGenreNotFound, err
	}

	paramsCheck := database.CheckBookGenreExistsParams{
		BookID:  int32(bookId),
		GenreID: int32(genreId),
	}
	if isExists, err := q.CheckBookGenreExists(ctx, paramsCheck); err != nil || isExists {
		return nil, response.CodeBookGenreExists, err
	}

	paramsInsert := database.InsertBookGenreParams{
		BookID:  int32(bookId),
		GenreID: int32(genreId),
	}

	result, err := q.InsertBookGenre(ctx, paramsInsert)
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeCannotInsertGenreToBook, err
	}
	_ = tx.Commit()
	return &result, response.CodeSuccess, nil
}
func (bs *bookService) AddAuthorToBook(ctx context.Context, bookId int, authorId int) (*database.BookAuthor, int, error) {
	tx, _ := bs.db.BeginTx(ctx, nil)
	q := bs.repository.WithTx(tx)

	if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
		return nil, response.CodeBookNotFound, err
	}

	if _, err := q.GetGenreByID(ctx, int32(authorId)); err != nil {
		return nil, response.CodeAuthorNotFound, err
	}

	paramsCheck := database.CheckBookAuthorExistsParams{
		BookID:   int32(bookId),
		AuthorID: int32(authorId),
	}
	if isExists, err := q.CheckBookAuthorExists(ctx, paramsCheck); err != nil || isExists {
		return nil, response.CodeBookAuthorExists, err
	}

	paramsInsert := database.InsertBookAuthorParams{
		BookID:   int32(bookId),
		AuthorID: int32(authorId),
	}

	result, err := q.InsertBookAuthor(ctx, paramsInsert)
	if err != nil {
		_ = tx.Rollback()
		return nil, response.CodeCannotInsertAuthorToBook, err
	}
	_ = tx.Commit()
	return &result, response.CodeSuccess, nil
}
