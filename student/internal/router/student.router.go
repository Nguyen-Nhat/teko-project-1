package router

import (
	"github.com/gin-gonic/gin"
	"student/internal/wire"
)

type StudentRouter struct {
}

func (ur *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	studentController := wire.InitStudentRouterHandler()
	studentRouter := Router.Group("/student")
	{
		studentRouter.GET("", studentController.GetPageStudentWithFilter)
		studentRouter.GET("/:student_id", studentController.GetStudentById)
		studentRouter.POST("", studentController.CreateStudent)
	}
}
