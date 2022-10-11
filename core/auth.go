package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Check auth")
		//user := 1
		//if user == 1 {
		//	APIResponse(c, nil, http.StatusUnauthorized, "Not auth")
		//}
		c.Next()
	}
}
