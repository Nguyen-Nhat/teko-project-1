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

// Create Book documentation
// @Summary      Create book
// @Description  Create book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        payload body req.BookPostDto true "payload"
// @Success      201  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Router       /book [post]
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
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.Id)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}

// Add Genre To Book documentation
// @Summary      Add Genre To Book
// @Description  Add Genre To Book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book_id   path      int  true  "Book Id"
// @Param        genre_id   path      int  true  "Genre Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /book/{book_id}/add-genre/{genre_id} [put]
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

// Add Author To Book documentation
// @Summary      Add Author To Book
// @Description  Add Author To Book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book_id   path      int  true  "Book Id"
// @Param        author_id   path      int  true  "Author Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /book/{book_id}/add-author/{author_id} [put]
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

// Remove Genre From Book documentation
// @Summary      Remove Genre From Book
// @Description  Remove Genre From Book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book_id   path      int  true  "Book Id"
// @Param        genre_id   path      int  true  "Genre Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /book/{book_id}/remove-genre/{genre_id} [put]
func (bc *BookController) RemoveGenreFromBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	genreIdStr := ctx.Param("genre_id")
	genreId, errGenreId := strconv.Atoi(genreIdStr)
	if errBookId != nil || errGenreId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	_, code, _ := bc.bookService.RemoveAuthorFromBook(ctx, bookId, genreId)
	response.Response(ctx, code, nil)
}

// Remove Author From Book documentation
// @Summary      Remove Author From Book
// @Description  Remove Author From Book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book_id   path      int  true  "Book Id"
// @Param        author_id   path      int  true  "Author Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /book/{book_id}/remove-author/{author_id} [put]
func (bc *BookController) RemoveAuthorFromBook(ctx *gin.Context) {
	bookIdStr := ctx.Param("book_id")
	bookId, errBookId := strconv.Atoi(bookIdStr)

	authorIdStr := ctx.Param("author_id")
	authorId, errAuthorId := strconv.Atoi(authorIdStr)
	if errBookId != nil || errAuthorId != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	_, code, _ := bc.bookService.RemoveAuthorFromBook(ctx, bookId, authorId)
	response.Response(ctx, code, nil)
}

// Get Book Detail By Id documentation
// @Summary      Get Book Detail By Id
// @Description  Get Book Detail By Id
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book_id   path      int  true  "Book Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /book/{book_id} [get]
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

// Get Page Book With Filter documentation
// @Summary      Get Page Book With Filter
// @Description  Get Page Book With Filter
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        query   query      req.BookPageDto  true  "Query Param"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /book [get]
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
