package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jinzhu/copier"
	"library/global"
	"library/internal/database"
	"library/internal/dto/req"
	"library/internal/dto/res"
	"library/internal/util"
	"library/pkg/response"
)

type IBookService interface {
	GetBookDetailById(ctx context.Context, bookId int) (*res.BookDetailDto, int, error)
	GetPageBookWithFilter(ctx context.Context, data *req.BookPageDto) (*res.PageResult, int, error)
	CreateBook(ctx context.Context, data *req.BookPostDto) (*database.Book, int, error)
	AddGenreToBook(ctx context.Context, bookId int, genreId int) (*database.BookGenre, int, error)
	AddAuthorToBook(ctx context.Context, bookId int, authorId int) (*database.BookAuthor, int, error)
	RemoveAuthorFromBook(ctx context.Context, bookId int, authorId int) (interface{}, int, error)
	RemoveGenreFromBook(ctx context.Context, bookId int, genreId int) (interface{}, int, error)
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
func (bs *bookService) GetPageBookWithFilter(ctx context.Context, data *req.BookPageDto) (*res.PageResult, int, error) {
	params := database.GetPageBookWithFilterParams{
		Title:    util.ToNullString(data.Title),
		Genreid:  util.ToNullInt32(data.GenreId),
		Authorid: util.ToNullInt32(data.AuthorId),
		Page:     data.Page,
		Size:     data.Size,
	}
	books, err := bs.repository.GetPageBookWithFilter(ctx, params)
	if err != nil {
		fmt.Println(err)
		return nil, response.CodeInternalServerError, err
	}
	totalPage := 0
	if len(books) > 0 {
		totalPage = int(books[0].TotalPage)
	}

	list := make([]res.BookDto, len(books))
	for i, book := range books {
		var b database.Book
		_ = copier.Copy(&b, &book)
		temp := res.BookDto{}
		temp.FromModel(b)
		list[i] = temp
	}
	result := res.PageResult{
		List:      list,
		Size:      len(list),
		Page:      int(data.Page),
		TotalPage: totalPage,
	}

	return &result, response.CodeSuccess, nil
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
	return withTransaction(ctx, bs.db, bs.repository, func(q *database.Queries) (*database.Book, int, error) {
		params := database.CreateBookParams{
			Title:         data.Title,
			PublishedYear: data.PublishedYear,
		}
		result, err := q.CreateBook(ctx, params)
		if err != nil {
			return nil, response.CodeBookNotFound, err
		}
		return &result, response.CodeSuccess, nil
	})
}
func (bs *bookService) AddGenreToBook(ctx context.Context, bookId int, genreId int) (*database.BookGenre, int, error) {
	return withTransaction(ctx, bs.db, bs.repository, func(q *database.Queries) (*database.BookGenre, int, error) {
		if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
			return nil, response.CodeBookNotFound, err
		}
		if _, err := q.GetGenreByID(ctx, int32(genreId)); err != nil {
			return nil, response.CodeGenreNotFound, err
		}
		paramsInsert := database.InsertBookGenreParams{
			BookID:  int32(bookId),
			GenreID: int32(genreId),
		}
		result, err := q.InsertBookGenre(ctx, paramsInsert)
		if err != nil {
			return nil, response.CodeCannotInsertGenreToBook, err
		}
		return &result, response.CodeSuccess, nil
	})
}
func (bs *bookService) AddAuthorToBook(ctx context.Context, bookId int, authorId int) (*database.BookAuthor, int, error) {
	return withTransaction(ctx, bs.db, bs.repository, func(q *database.Queries) (*database.BookAuthor, int, error) {
		if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
			return nil, response.CodeBookNotFound, err
		}

		if _, err := q.GetGenreByID(ctx, int32(authorId)); err != nil {
			return nil, response.CodeAuthorNotFound, err
		}

		paramsInsert := database.InsertBookAuthorParams{
			BookID:   int32(bookId),
			AuthorID: int32(authorId),
		}

		result, err := q.InsertBookAuthor(ctx, paramsInsert)
		if err != nil {
			return nil, response.CodeCannotInsertAuthorToBook, err
		}
		return &result, response.CodeSuccess, nil
	})
}

func (bs *bookService) RemoveAuthorFromBook(ctx context.Context, bookId int, authorId int) (interface{}, int, error) {
	return withTransaction(ctx, bs.db, bs.repository, func(q *database.Queries) (interface{}, int, error) {
		if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
			return nil, response.CodeBookNotFound, err
		}
		if _, err := q.GetGenreByID(ctx, int32(authorId)); err != nil {
			return nil, response.CodeAuthorNotFound, err
		}
		paramsRemove := database.RemoveBookAuthorParams{
			BookID:   int32(bookId),
			AuthorID: int32(authorId),
		}
		err := q.RemoveBookAuthor(ctx, paramsRemove)
		if err != nil {
			return nil, response.CodeCannotRemoveAuthorFromBook, err
		}
		return nil, response.CodeSuccess, nil
	})
}
func (bs *bookService) RemoveGenreFromBook(ctx context.Context, bookId int, genreId int) (interface{}, int, error) {
	return withTransaction(ctx, bs.db, bs.repository, func(q *database.Queries) (interface{}, int, error) {
		if _, err := q.GetBookByID(ctx, int32(bookId)); err != nil {
			return nil, response.CodeBookNotFound, err
		}

		if _, err := q.GetGenreByID(ctx, int32(genreId)); err != nil {
			return nil, response.CodeGenreNotFound, err
		}

		paramsRemove := database.RemoveBookGenreParams{
			BookID:  int32(bookId),
			GenreID: int32(genreId),
		}
		if err := q.RemoveBookGenre(ctx, paramsRemove); err != nil {
			return nil, response.CodeCannotRemoveGenreFromBook, err
		}
		return nil, response.CodeSuccess, nil
	})
}
