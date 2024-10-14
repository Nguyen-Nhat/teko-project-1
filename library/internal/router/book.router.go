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
		bookRouter.POST("", bookController.CreateBook)
		bookRouter.PUT("/:book_id/add-genre/:genre_id", bookController.AddGenreToBook)
		bookRouter.PUT("/:book_id/add-genre/:author_id", bookController.AddAuthorToBook)
	}
}
