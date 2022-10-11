package auth

import "github.com/gin-gonic/gin"

type AuthController struct {
}

func (receiver AuthController) Login(c *gin.Context) {
	c.JSON(200, "login")
}
