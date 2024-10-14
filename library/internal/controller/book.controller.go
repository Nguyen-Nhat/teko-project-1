package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/request"
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
	var book request.BookPostDto
	if err := ctx.ShouldBindJSON(&book); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
	}
	result, code, err := bc.bookService.CreateBook(ctx, &book)
	if err != nil {
		response.Response(ctx, code, nil)
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
	}
	result, code, _ := bc.bookService.AddGenreToBook(ctx, bookId, genreId)
	response.Response(ctx, code, result)
}
func (bc *BookController) AddAuthorToBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	authorIdStr := ctx.Param("genre_id")
	authorId, errAuthorId := strconv.Atoi(authorIdStr)
	if errBookId != nil || errAuthorId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
	}
	result, code, _ := bc.bookService.AddAuthorToBook(ctx, bookId, authorId)
	response.Response(ctx, code, result)
}
