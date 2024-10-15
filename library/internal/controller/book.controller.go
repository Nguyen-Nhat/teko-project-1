package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/req"
	"library/internal/service"
	"library/pkg/response"
	"strconv"
)

type BookController struct {
	bookService service.IBookService
}

func NewBookController(bookService service.IBookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var book req.BookPostDto
	if err := ctx.ShouldBindJSON(&book); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, err := bc.bookService.CreateBook(ctx, &book)
	if err != nil {
		response.Response(ctx, code, nil)
		return
	}
	baseUrl := ctx.Request.URL.Path
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.ID)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}
func (bc *BookController) AddGenreToBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	genreIdStr := ctx.Param("genre_id")
	genreId, errGenreId := strconv.Atoi(genreIdStr)
	if errBookId != nil || errGenreId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := bc.bookService.AddGenreToBook(ctx, bookId, genreId)
	response.Response(ctx, code, result)
}
func (bc *BookController) AddAuthorToBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	authorIdStr := ctx.Param("author_id")
	authorId, errAuthorId := strconv.Atoi(authorIdStr)
	if errBookId != nil || errAuthorId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := bc.bookService.AddAuthorToBook(ctx, bookId, authorId)
	response.Response(ctx, code, result)
}
func (bc *BookController) RemoveGenreFromBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	genreIdStr := ctx.Param("genre_id")
	genreId, errGenreId := strconv.Atoi(genreIdStr)
	if errBookId != nil || errGenreId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	code, _ := bc.bookService.RemoveAuthorFromBook(ctx, bookId, genreId)
	response.Response(ctx, code, nil)
}
func (bc *BookController) RemoveAuthorFromBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	authorIdStr := ctx.Param("author_id")
	authorId, errAuthorId := strconv.Atoi(authorIdStr)
	if errBookId != nil || errAuthorId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	code, _ := bc.bookService.RemoveAuthorFromBook(ctx, bookId, authorId)
	response.Response(ctx, code, nil)
}
func (bc *BookController) GetBookDetailById(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)
	if errBookId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := bc.bookService.GetBookDetailById(ctx, bookId)
	response.Response(ctx, code, result)
}
func (bc *BookController) GetPageBookWithFilter(ctx *gin.Context) {
	var query req.BookPageDto
	query.SetDefaultPageInfo()
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Response(ctx, response.CodeInvalidRequestParam, nil)
		return
	}
	result, code, _ := bc.bookService.GetPageBookWithFilter(ctx, &query)
	response.Response(ctx, code, result)
}
