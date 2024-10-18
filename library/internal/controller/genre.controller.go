package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/internal/dto/req"
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

// Create Genre documentation
// @Summary      Create genre
// @Description  Create genre
// @Tags         genre
// @Accept       json
// @Produce      json
// @Param        payload body req.GenrePostDto true "payload"
// @Success      201  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /genre [post]
func (gc *GenreController) CreateGenre(ctx *gin.Context) {
	var genre req.GenrePostDto
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
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.Id)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}

// Get Genre By Id documentation
// @Summary      Get Genre By Id
// @Description  Get Genre By Id
// @Tags         genre
// @Accept       json
// @Produce      json
// @Param        genre_id   path      int  true  "Genre Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Router       /genre/{genre_id} [get]
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
