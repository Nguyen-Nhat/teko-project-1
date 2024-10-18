package controller

import (
	"github.com/gin-gonic/gin"
	"library/internal/dto/req"
	"library/internal/service"
	"library/pkg/response"
	"strconv"
)

type BorrowBookController struct {
	borrowBookService service.IBorrowBookService
}

func NewBorrowBookController(borrowBookService service.IBorrowBookService) *BorrowBookController {
	return &BorrowBookController{
		borrowBookService: borrowBookService,
	}
}
func (bbc *BorrowBookController) GetBorrowBookDetails(ctx *gin.Context) {
	var query req.BorrowBookDetailPageDto
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, _ := bbc.borrowBookService.GetBorrowBookDetails(ctx, &query)
	response.Response(ctx, code, result)
}
func (bbc *BorrowBookController) CreateBorrowBook(ctx *gin.Context) {
	var borrowBook req.BorrowBookPostDto
	if err := ctx.ShouldBindJSON(&borrowBook); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, _ := bbc.borrowBookService.CreateBorrowBook(ctx, &borrowBook)
	response.Response(ctx, code, result)
}

func (bbc *BorrowBookController) ReturnBorrowBook(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := bbc.borrowBookService.ReturnBorrowBook(ctx, id)
	response.Response(ctx, code, result)
}
