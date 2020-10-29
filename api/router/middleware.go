package router

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func IsAdmin(context *gin.Context) {
	claims := jwt.ExtractClaims(context)

	role := claims["role"].(string)
	if role == "admin" {
		context.Next()
		return
	}
	context.JSON(401, gin.H{"error": "You don't have access to this route"})
}
