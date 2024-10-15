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
		bookRouter.GET("", bookController.GetPageBookWithFilter)
		bookRouter.GET("/:book_id", bookController.GetBookDetailById)
		bookRouter.POST("", bookController.CreateBook)
		bookRouter.PUT("/:book_id/add-genre/:genre_id", bookController.AddGenreToBook)
		bookRouter.PUT("/:book_id/add-author/:author_id", bookController.AddAuthorToBook)
		bookRouter.PUT("/:book_id/remove-genre/:genre_id", bookController.RemoveGenreFromBook)
		bookRouter.PUT("/:book_id/remove-author/:author_id", bookController.RemoveAuthorFromBook)
	}
}
