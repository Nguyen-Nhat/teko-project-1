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

// Get Borrow Book Details documentation
// @Summary      Get Borrow Book Details
// @Description  Get Borrow Book Details
// @Tags         borrow book
// @Accept       json
// @Produce      json
// @Param        query query req.BorrowBookDetailPageDto true "query"
// @Success      200  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /borrow-book/detail-borrow-book [get]
func (bbc *BorrowBookController) GetBorrowBookDetails(ctx *gin.Context) {
	var query req.BorrowBookDetailPageDto
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, _ := bbc.borrowBookService.GetBorrowBookDetails(ctx, &query)
	response.Response(ctx, code, result)
}

// Create Borrow Book documentation
// @Summary      Create borrow book
// @Description  Create borrow book
// @Tags         borrow book
// @Accept       json
// @Produce      json
// @Param        payload body req.BorrowBookPostDto true "payload"
// @Success      201  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /borrow-book [post]
func (bbc *BorrowBookController) CreateBorrowBook(ctx *gin.Context) {
	var borrowBook req.BorrowBookPostDto
	if err := ctx.ShouldBindJSON(&borrowBook); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, _ := bbc.borrowBookService.CreateBorrowBook(ctx, &borrowBook)
	response.Response(ctx, code, result)
}

// Return Borrow Book documentation
// @Summary      Return Borrow Book
// @Description  Return Borrow Book
// @Tags         borrow book
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Borrow Book Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /borrow-book/return/{id} [put]
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
