package auth

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("v1/auth")
	{
		auth.GET("/login", AuthController{}.Login)
	}
}
