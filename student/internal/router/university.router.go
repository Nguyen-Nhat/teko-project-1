package router

import (
	"github.com/gin-gonic/gin"
)

type UniversityRouter struct {
}

func (ur *UniversityRouter) InitUniversityRouter(Router *gin.RouterGroup) {
	//universityController := wire.InitUniversityRouterHandler()
	universityRouter := Router.Group("/book")
	{
		universityRouter.GET("/test1")
		universityRouter.POST("/test2")
	}
}
