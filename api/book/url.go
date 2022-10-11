package book

import (
	"github.com/gin-gonic/gin"
	"go-gin/core"
)

func ApplyRoutes(r *gin.RouterGroup) {
	book := r.Group("v1/book", core.AuthMiddleware())
	{
		book.GET("/", bookApiController{}.GetBook)
		book.POST("/", bookApiController{}.CreateBook)
		book.GET("/:id", bookApiController{}.RetrieveBook)
		book.PUT("/:id", bookApiController{}.UpdateBook)
		book.DELETE("/:id", bookApiController{}.DeleteBook)
	}
}
