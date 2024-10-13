package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"student/global"
)

func LimitRequestBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.ContentLength > global.Config.Server.MaxRequestBody {
			c.AbortWithStatus(http.StatusRequestEntityTooLarge)
			return
		}
		c.Next()
	}
}
