package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/request"
	"library/internal/service"
	"library/pkg/response"
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
	}
	result, code, err := ac.authorService.CreateAuthor(ctx, &author)
	if err != nil {
		response.Response(ctx, code, nil)
	}
	baseUrl := ctx.Request.URL.Path
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.ID)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}
