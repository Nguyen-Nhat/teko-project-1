package middleware

import (
	"github.com/gin-gonic/gin"
	"library/global"
	"library/pkg/setting"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	if global.Config.Cors.Mode == "all" {
		return defaultCors()
	}
	return func(c *gin.Context) {
		whitelist := checkCors(c.GetHeader("origin"))

		if whitelist != nil {
			c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
			c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
			c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
			c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
			if whitelist.AllowCredentials {
				c.Header("Access-Control-Allow-Credentials", "true")
			}
			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusNoContent)
			}
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}

func defaultCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
func checkCors(currentOrigin string) *setting.CORSWhitelist {
	for _, whitelist := range global.Config.Cors.WhiteList {
		if currentOrigin == whitelist.AllowOrigin {
			return &whitelist
		}
	}
	return nil
}
