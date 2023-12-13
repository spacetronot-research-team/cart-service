package response

import "github.com/gin-gonic/gin"

var statusMessage map[statusCode]string = map[statusCode]string{
	"CART2001000": "success",
	"CART5001000": "unknown error",
	"CART5001001": "error bar karena blablabla",
}

type statusCode string

func Write(ctx *gin.Context, httpCode int, data interface{}, statusCode statusCode) {
	body := gin.H{
		"data":       data,
		"statusCode": statusCode,
	}
	if message, ok := statusMessage[statusCode]; ok {
		body["message"] = message
	}
	ctx.JSON(httpCode, body)
}
