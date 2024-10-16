package router

import (
	"github.com/gin-gonic/gin"
	"library/internal/wire"
)

type BorrowBookRouter struct {
}

func (br *BorrowBookRouter) InitBorrowBookRouter(Router *gin.RouterGroup) {
	borrowBookController := wire.InitBorrowBookRouterHandler()
	borrowBookRouter := Router.Group("/borrow-book")
	{
		borrowBookRouter.GET("")
		borrowBookRouter.POST("", borrowBookController.CreateBorrowBook)
		borrowBookRouter.PUT("/:id", borrowBookController.ReturnBorrowBook)
	}
}
