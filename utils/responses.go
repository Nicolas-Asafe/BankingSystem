package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string
	Error   error
	Data    any
}

func Resthis(Message string, Error error, Data any) Response {
	return Response{Message, Error, Data}
}

func ReciveForLog(response Response, logs *Logs) {
	if response.Error != nil {
		logs.Error(response.Error.Error())
		return
	}
	logs.Info(response.Message)
}

func ReciveForGin(response Response, status int, ctx *gin.Context) {
	ctx.JSON(status,
		gin.H{
			"message": response.Message,
			"error":   response.Error.Error(),
			"data":    response.Data,
		})
}
