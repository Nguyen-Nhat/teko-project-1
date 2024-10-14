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
		genreRouter.POST("", genreController.CreateGenre)
	}
}
