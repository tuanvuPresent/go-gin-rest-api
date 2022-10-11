package book

import (
	"github.com/gin-gonic/gin"
	"go-gin/core"
)

func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("v1/book", core.AuthMiddleware())
	{
		auth.GET("/", bookApiController{}.GetBook)
		auth.POST("/", bookApiController{}.CreateBook)
		auth.GET("/:id", bookApiController{}.RetrieveBook)
		auth.PUT("/:id", bookApiController{}.UpdateBook)
		auth.DELETE("/:id", bookApiController{}.DeleteBook)
	}
}
