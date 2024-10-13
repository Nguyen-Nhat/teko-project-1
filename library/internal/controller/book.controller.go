package controller

import (
	"github.com/gin-gonic/gin"
	"library/internal/service"
	"library/pkg/response"
)

type BookController struct {
	bookService service.IBookService
}

func NewBookController(bookService service.IBookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}
func (bc *BookController) GetAll(ctx *gin.Context) {
	result, _ := bc.bookService.GetAll(ctx)
	response.SuccessResponse(ctx, response.CodeSuccess, result)
}
