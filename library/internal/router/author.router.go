package router

import (
	"github.com/gin-gonic/gin"
	"library/internal/wire"
)

type AuthorRouter struct {
}

func (br *AuthorRouter) InitAuthorRouter(Router *gin.RouterGroup) {
	authorController := wire.InitAuthorRouterHandler()
	authorRouter := Router.Group("author")
	{
		authorRouter.GET("/:author_id", authorController.GetAuthorById)
		authorRouter.POST("", authorController.CreateAuthor)
	}

}
