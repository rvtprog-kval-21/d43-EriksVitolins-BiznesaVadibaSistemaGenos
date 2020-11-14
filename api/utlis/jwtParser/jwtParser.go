package jwtParser

import (
	"api/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetClaims(context *gin.Context) jwt.MapClaims {
	requestToken := context.Request.Header["Authorization"][0]
	requestToken = strings.Split(requestToken, " ")[1]
	secretString := config.Secret
	secret := []byte(secretString)
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return secret, nil
	})

	if err != nil {
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims

	}
	return nil
}
