package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var TraceID = "X-Trace-Id"

func TraceIDMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		var traceID string = uuid.New().String()
		context.Header(TraceID, traceID)
		context.Set(TraceID, traceID)
		context.Next()
	}
}
