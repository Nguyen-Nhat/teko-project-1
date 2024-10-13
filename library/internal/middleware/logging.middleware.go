package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Record struct {
	Ip           string        `json:"ip"`
	Method       string        `json:"method"`
	Path         string        `json:"path"`
	Status       int           `json:"status"`
	Latency      time.Duration `json:"latency"`
	Agent        string        `json:"agent"`
	RequestBody  string        `json:"request_body"`
	ResponseBody string        `json:"response_body"`
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//body, err := io.ReadAll(c.Request.Body)
		//now := time.Now()
		//c.Next()
		//record := Record{
		//	Ip:      c.ClientIP(),
		//	Method:  c.Request.Method,
		//	Path:    c.Request.URL.Path,
		//	Agent:   c.Request.UserAgent(),
		//	Latency: time.Since(now),
		//	Status:  c.Writer.Status(),
		//	ResponseBody: c.Request.Body
		//}
	}
}
