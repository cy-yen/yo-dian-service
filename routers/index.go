package routers

import (
	"yo-dian-service/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")

	{
		user.GET("/ping", (&controllers.UserController{}).Get)
	}

	return r
}
