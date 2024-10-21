package router

import (
	"github.com/gin-gonic/gin"
	"student/internal/wire"
)

type UniversityRouter struct {
}

func (ur *UniversityRouter) InitUniversityRouter(Router *gin.RouterGroup) {
	universityController := wire.InitUniversityRouterHandler()
	universityRouter := Router.Group("/book")
	{
		universityRouter.GET("/:university_id", universityController.GetUniversityById)
		universityRouter.POST("", universityController.CreateUniversity)
	}
}
