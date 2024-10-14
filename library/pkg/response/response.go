package response

import (
	"github.com/gin-gonic/gin"
)

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *gin.Context, code int, data interface{}) {
	c.JSON(msg[code].HttpCode, Data{
		Code:    code,
		Message: msg[code].Message,
		Data:    data,
	})
}
