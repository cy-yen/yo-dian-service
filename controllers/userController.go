package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"yo-dian-service/common"
)

type UserController struct{}

func (u *UserController) Get(c *gin.Context) {
	common.SuccessResponse(c, http.StatusOK, "success", "pong pong")
}
