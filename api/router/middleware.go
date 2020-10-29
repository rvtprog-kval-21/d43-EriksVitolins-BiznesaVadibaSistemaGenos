package router

import (
	"api/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

type Tokens struct {
	Token string `form:"token"`
}

func IsAdmin(context *gin.Context) {
	requestToken := context.Request.Header["Authorization"][0]
	requestToken = strings.Split(requestToken," ")[1]
	secretString := config.Secret
	secret := []byte(secretString)
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return secret, nil
	})

	if err != nil {
		context.JSON(401, gin.H{"error": "You don't have access to this route"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role"] == "admin" {
			context.Next()
			return
		}
	}
	context.JSON(401, gin.H{"error": "You don't have access to this route"})
}
