package router

import (
	"github.com/gin-gonic/gin"
	"library/internal/wire"
)

type BookRouter struct {
}

func (br *BookRouter) InitBookRouter(Router *gin.RouterGroup) {
	bookController := wire.InitBookRouterHandler()
	bookRouter := Router.Group("/book")
	{
		bookRouter.GET("/test1", bookController.GetAll)
		bookRouter.POST("/test2")
	}
}
