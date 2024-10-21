package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"student/internal/dto/req"
	"student/internal/service"
	"student/pkg/response"
)

type StudentController struct {
	studentService service.IStudentService
}

func NewStudentController(studentService service.IStudentService) *StudentController {
	return &StudentController{
		studentService: studentService,
	}
}

// Create Student documentation
// @Summary      Create student
// @Description  Create student
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        payload body req.StudentPostDto true "payload"
// @Success      201  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /student [post]
func (sc *StudentController) CreateStudent(ctx *gin.Context) {
	var body req.StudentPostDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Response(ctx, response.CodeInvalidRequestBody, nil)
		return
	}
	result, code, err := sc.studentService.CreateStudent(ctx, &body)
	if err != nil {
		response.Response(ctx, code, nil)
		return
	}
	baseUrl := ctx.Request.URL.Path
	locationUrl := fmt.Sprintf("%s/%d", baseUrl, result.Id)
	ctx.Header("Location", locationUrl)
	response.Response(ctx, code, result)
}

// Get Student By Id documentation
// @Summary      Get Student By Id
// @Description  Get Student By Id
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        student_id   path      int  true  "Student Id"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      404  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /student/{student_id} [get]
func (sc *StudentController) GetStudentById(ctx *gin.Context) {
	studentIdStr := ctx.Param("student_id")
	studentId, err := strconv.Atoi(studentIdStr)
	if err != nil || studentId <= 0 {
		response.Response(ctx, response.CodeInvalidPathVariable, nil)
		return
	}
	result, code, _ := sc.studentService.GetStudentById(ctx, studentId)
	response.Response(ctx, code, result)
}

// Get Student With Filter documentation
// @Summary      Get Student With Filter
// @Description  Get Student With Filter
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        query   query      req.StudentPageDto  true  "Query Param"
// @Success      200  {object}  response.Data
// @Failure      400  {object}  response.Data
// @Failure      500  {object}  response.Data
// @Router       /student [get]
func (sc *StudentController) GetPageStudentWithFilter(ctx *gin.Context) {
	var query req.StudentPageDto
	query.SetDefaultPageInfo()
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Response(ctx, response.CodeInvalidRequestParam, nil)
		return
	}
	result, code, _ := sc.studentService.GetPageStudentWithFilter(ctx, &query)
	response.Response(ctx, code, result)
}
