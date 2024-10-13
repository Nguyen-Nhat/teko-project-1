package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDate struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse 200
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// CreatedResponse 201
func CreatedResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusCreated, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

// BadRequestResponse 400
func BadRequestResponse(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// UnauthorizedResponse 401
func UnauthorizedResponse(c *gin.Context, code int) {
	c.JSON(http.StatusUnauthorized, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// ForbiddenResponse 403
func ForbiddenResponse(c *gin.Context, code int) {
	c.JSON(http.StatusForbidden, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// NotFoundResponse 404
func NotFoundResponse(c *gin.Context, code int) {
	c.JSON(http.StatusNotFound, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}

// InternalServerErrorResponse 500
func InternalServerErrorResponse(c *gin.Context, code int) {
	c.JSON(http.StatusInternalServerError, ResponseDate{
		Code:    code,
		Message: msg[code],
		Data:    nil,
	})
}
