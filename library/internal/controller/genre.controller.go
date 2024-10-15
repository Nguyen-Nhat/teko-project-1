package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/request"
	"library/internal/service"
	"library/pkg/response"
	"strconv"
)

type GenreController struct {
	genreService service.IGenreService
}

func NewGenreController(genreService service.IGenreService) *GenreController {
	return &GenreController{
		genreService: genreService,
	}
}
func (gc *GenreController) CreateGenre(ctx *gin.Context) {
	var genre request.GenrePostDto
	if err := ctx.ShouldBindJSON(&genre); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, err := gc.genreService.CreateGenre(ctx, &genre)
	if err != nil {
		response.Response(ctx, code, nil)
		return
	}
	baseUrl := ctx.Request.URL.Path
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.ID)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}
func (gc *GenreController) GetGenreById(ctx *gin.Context) {
	genreIdStr := ctx.Param("genre_id")
	genreId, err := strconv.Atoi(genreIdStr)
	if err == nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := gc.genreService.GetGenreById(ctx, genreId)
	response.Response(ctx, code, result)
}
