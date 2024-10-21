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
func (uc *UniversityController) GetUniversityById(ctx *gin.Context) {
	universityIdStr := ctx.Param("university_id")
	universityId, err := strconv.Atoi(universityIdStr)
	if err != nil {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := uc.universityService.GetUniversityById(ctx, universityId)
	response.Response(ctx, code, result)
}
