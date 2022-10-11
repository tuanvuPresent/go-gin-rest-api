package core

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func APIResponse(ctx *gin.Context, Data interface{}, StatusCode int, Message string) {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	}
	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}
