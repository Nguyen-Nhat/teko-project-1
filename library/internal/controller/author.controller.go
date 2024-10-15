package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/request"
	"library/internal/service"
	"library/pkg/response"
	"strconv"
)

type AuthorController struct {
	authorService service.IAuthorService
}

func NewAuthorController(authorService service.IAuthorService) *AuthorController {
	return &AuthorController{
		authorService: authorService,
	}
}

func (ac *AuthorController) CreateAuthor(ctx *gin.Context) {
	var author request.AuthorPostDto
	if err := ctx.ShouldBindJSON(&author); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, err := ac.authorService.CreateAuthor(ctx, &author)
	if err != nil {
		response.Response(ctx, code, nil)
		return
	}
	baseUrl := ctx.Request.URL.Path
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.ID)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}
func (ac *AuthorController) GetAuthorById(ctx *gin.Context) {
	authorIdStr := ctx.Param("author_id")
	authorId, err := strconv.Atoi(authorIdStr)
	if err != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := ac.authorService.GetAuthorById(ctx, authorId)
	response.Response(ctx, code, result)
}
