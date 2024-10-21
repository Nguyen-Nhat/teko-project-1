package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"student/internal/dto/req"
	"student/internal/service"
	"student/pkg/response"
)

type UniversityController struct {
	universityService service.IUniversityService
}

func NewUniversityController(universityService service.IUniversityService) *UniversityController {
	return &UniversityController{
		universityService: universityService,
	}
}

// Create University documentation
// @Summary      Create university
// @Description  Create university
// @Tags         university
// @Accept       json
// @Produce      json
// @Param        payload body req.UniversityPostDto true "payload"
// @Success      201  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /university [post]
func (uc *UniversityController) CreateUniversity(ctx *gin.Context) {
	var body req.UniversityPostDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, err := uc.universityService.CreateUniversity(ctx, &body)
	if err != nil {
		response.Response(ctx, code, nil)
		return
	}
	baseUrl := ctx.Request.URL.Path
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.Id)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}

// Get University By Id documentation
// @Summary      Get University By Id
// @Description  Get University By Id
// @Tags         university
// @Accept       json
// @Produce      json
// @Param        university_id   path      int  true  "University Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /university/{university_id} [get]
func (uc *UniversityController) GetUniversityById(ctx *gin.Context) {
	universityIdStr := ctx.Param("university_id")
	universityId, err := strconv.Atoi(universityIdStr)
	if err != nil || universityId <= 0 {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := uc.universityService.GetUniversityById(ctx, universityId)
	response.Response(ctx, code, result)
}
