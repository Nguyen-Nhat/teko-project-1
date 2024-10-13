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
		studentRouter.GET("/test1")
		studentRouter.POST("/test2")
	}
}
