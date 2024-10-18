package router

import (
	"github.com/gin-gonic/gin"
)

type StudentRouter struct {
}

func (ur *StudentRouter) InitStudentRouter(Router *gin.RouterGroup) {
	//studentController := wire.InitStudentRouterHandler()
	studentRouter := Router.Group("/book")
	{
		studentRouter.GET("")
		studentRouter.GET("/:student_id")
		studentRouter.POST("")
	}
}
