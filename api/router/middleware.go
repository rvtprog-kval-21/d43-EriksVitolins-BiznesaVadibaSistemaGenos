package router

import (
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
)

type Tokens struct {
	Token string `form:"token"`
}

func IsAdmin(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims != nil {
		if claims["role"] == "admin" {
			context.Next()
			return
		}
	}
	context.JSON(401, gin.H{"error": "You don't have access to this route"})
}
