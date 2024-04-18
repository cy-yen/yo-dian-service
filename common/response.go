package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Success struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(http.StatusOK, &Success{Code: statusCode, Message: message, Data: data})
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(http.StatusOK, &Error{Code: statusCode, Message: message})
}
