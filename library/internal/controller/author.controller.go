package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/req"
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

// Create Author documentation
// @Summary      Create author
// @Description  Create author
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        payload body req.AuthorPostDto true "payload"
// @Success      201  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /author [post]
func (ac *AuthorController) CreateAuthor(ctx *gin.Context) {
	var author req.AuthorPostDto
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

// Get Author By Id documentation
// @Summary      Get Author By Id
// @Description  Get Author By Id
// @Tags         author
// @Accept       json
// @Produce      json
// @Param        author_id   path      int  true  "Author Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /author/{author_id} [get]
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
