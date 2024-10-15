package service

import (
	"context"
	"database/sql"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/internal/dto/res"
	"library/pkg/response"
)

type IBookService interface {
	GetBookDetailById(ctx context.Context, bookId int) (*res.BookDetailDto, int, error)
	GetPageBookWithFilter(ctx context.Context, data *req.BookPageDto) (interface{}, int, error)
	CreateBook(ctx context.Context, data *req.BookPostDto) (*database.Book, int, error)
	AddGenreToBook(ctx context.Context, bookId int, genreId int) (*database.BookGenre, int, error)
	AddAuthorToBook(ctx context.Context, bookId int, authorId int) (*database.BookAuthor, int, error)
	RemoveAuthorFromBook(ctx context.Context, bookId int, authorId int) (int, error)
	RemoveGenreFromBook(ctx context.Context, bookId int, genreId int) (int, error)
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
func (bs *bookService) GetPageBookWithFilter(ctx context.Context, data *req.BookPageDto) (interface{}, int, error) {
	return nil, response.CodeSuccess, nil
}
func (bs *bookService) GetBookDetailById(ctx context.Context, bookId int) (*res.BookDetailDto, int, error) {
	book, err := bs.repository.GetBookByID(ctx, int32(bookId))
	if err != nil {
		return nil, response.CodeBookNotFound, err
	}

	authors, err := bs.repository.GetAuthorsByBookID(ctx, int32(bookId))
	if err != nil {
		return nil, response.CodeInvalidPathVariable, err
	}

	genres, err := bs.repository.GetGenreByBookID(ctx, int32(bookId))
	if err != nil {
		return nil, response.CodeInvalidPathVariable, err
	}

	result := &res.BookDetailDto{}
	result.FromModel(book, authors, genres)
	return result, response.CodeSuccess, nil
}
func (bs *bookService) CreateBook(ctx context.Context, data *req.BookPostDto) (*database.Book, int, error) {
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
		_ = tx.Rollback()
		return nil, response.CodeBookNotFound, err
	}

	if _, err := q.GetGenreByID(ctx, int32(genreId)); err != nil {
		_ = tx.Rollback()
		return nil, response.CodeGenreNotFound, err
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
		_ = tx.Rollback()
		return nil, response.CodeBookNotFound, err
	}

	if _, err := q.GetGenreByID(ctx, int32(authorId)); err != nil {
		_ = tx.Rollback()
		return nil, response.CodeAuthorNotFound, err
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

func (bs *bookService) RemoveAuthorFromBook(ctx context.Context, bookId int, authorId int) (int, error) {
	tx, _ := bs.db.BeginTx(ctx, nil)
	q := bs.repository.WithTx(tx)

	if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
		_ = tx.Rollback()
		return response.CodeBookNotFound, err
	}

	if _, err := q.GetGenreByID(ctx, int32(authorId)); err != nil {
		_ = tx.Rollback()
		return response.CodeAuthorNotFound, err
	}

	paramsRemove := database.RemoveBookAuthorParams{
		BookID:   int32(bookId),
		AuthorID: int32(authorId),
	}

	err := q.RemoveBookAuthor(ctx, paramsRemove)
	if err != nil {
		_ = tx.Rollback()
		return response.CodeCannotRemoveAuthorFromBook, err
	}
	_ = tx.Commit()
	return response.CodeSuccess, nil
}
func (bs *bookService) RemoveGenreFromBook(ctx context.Context, bookId int, genreId int) (int, error) {
	tx, _ := bs.db.BeginTx(ctx, nil)
	q := bs.repository.WithTx(tx)

	if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
		_ = tx.Rollback()
		return response.CodeBookNotFound, err
	}

	if _, err := q.GetGenreByID(ctx, int32(genreId)); err != nil {
		_ = tx.Rollback()
		return response.CodeGenreNotFound, err
	}

	paramsRemove := database.RemoveBookGenreParams{
		BookID:  int32(bookId),
		GenreID: int32(genreId),
	}
	if err := q.RemoveBookGenre(ctx, paramsRemove); err != nil {
		_ = tx.Rollback()
		return response.CodeCannotRemoveGenreFromBook, err
	}
	_ = tx.Commit()
	return response.CodeSuccess, nil
}
