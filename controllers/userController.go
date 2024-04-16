package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "5555",
	})
}
