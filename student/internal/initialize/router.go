package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "student/docs"
	"student/global"
	"student/internal/middleware"
	"student/internal/router"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.LimitRequestBody())
	r.Use(middleware.CorsMiddleware())

	MainGroup := r.Group("/api/v1")
	{
		router.RouterGroupApp.InitStudentRouter(MainGroup)
		router.RouterGroupApp.InitUniversityRouter(MainGroup)
	}
	return r
}
