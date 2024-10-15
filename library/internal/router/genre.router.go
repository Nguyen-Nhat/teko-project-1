package router

import (
	"github.com/gin-gonic/gin"
	"library/internal/wire"
)

type GenreRouter struct {
}

func (br *GenreRouter) InitGenreRouter(Router *gin.RouterGroup) {
	genreController := wire.InitGenreRouterHandler()
	genreRouter := Router.Group("genre")
	{
		genreRouter.GET("/:genre_id", genreController.GetGenreById)
		genreRouter.POST("", genreController.CreateGenre)
	}
}
