package book

import (
	"github.com/gin-gonic/gin"
	"go-gin/core"
	"go-gin/database/models"
	"net/http"
)

type bookApiController struct {
	service bookService
}

func (receiver bookApiController) GetBook(c *gin.Context) {
	books := receiver.service.GetBook()

	core.APIResponse(c, books, http.StatusOK, "")
}

func (receiver bookApiController) CreateBook(c *gin.Context) {
	var requestBody RequestPostBook
	if err := c.BindJSON(&requestBody); err != nil {
		core.APIResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	book := models.Book{Description: requestBody.Description, Name: requestBody.Name}
	book, err := receiver.service.CreateBook(book)
	if err != nil {
		core.APIResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}

	core.APIResponse(c, book, http.StatusOK, "")
}

func (receiver bookApiController) RetrieveBook(c *gin.Context) {
	id := c.Param("id")
	book, err := receiver.service.RetrieveBook(id)
	if err != nil {
		core.APIResponse(c, nil, http.StatusNotFound, err.Error())
		return
	}

	core.APIResponse(c, book, http.StatusOK, "")
}

func (receiver bookApiController) UpdateBook(c *gin.Context) {
	var requestBody RequestPostBook
	if err := c.BindJSON(&requestBody); err != nil {
		core.APIResponse(c, nil, http.StatusBadRequest, err.Error())
		return
	}
	id := c.Param("id")

	book := models.Book{Description: requestBody.Description, Name: requestBody.Name}
	book, err := receiver.service.UpdateBook(id, book)
	if err != nil {
		core.APIResponse(c, nil, http.StatusNotFound, err.Error())
		return
	}

	core.APIResponse(c, book, http.StatusOK, "")
}

func (receiver bookApiController) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	book, err := receiver.service.DeleteBook(id)
	if err != nil {
		core.APIResponse(c, nil, http.StatusNotFound, err.Error())
		return
	}

	core.APIResponse(c, book, http.StatusOK, "")
}
