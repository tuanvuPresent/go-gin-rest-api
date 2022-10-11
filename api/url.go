package api

import (
	"github.com/gin-gonic/gin"
	"go-gin/api/auth"
	"go-gin/api/book"
)

func ApplyRoutes(r *gin.Engine) {
	v1 := r.Group("api/")
	{
		auth.ApplyRoutes(v1)
		book.ApplyRoutes(v1)
	}
}
