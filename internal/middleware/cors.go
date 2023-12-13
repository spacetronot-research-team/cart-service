package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCorsConfig() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"OPTIONS", "PUT", "POST", "GET", "DELETE"}
	return cors.New(corsConfig)
}
