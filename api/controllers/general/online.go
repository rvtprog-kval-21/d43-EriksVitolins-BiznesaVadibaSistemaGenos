package general

import (
	"api/model/user"
	"api/services/online"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	online.Ping(claims["id"])
	context.JSON(200, gin.H{})
	return
}

func UsersOnline(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	var keys []interface{}
	for key, _ := range online.Online {
		if key == claims["id"] {
			continue
		}
		keys = append(keys, key)
	}

	users := user.GetUsersIn(keys)

	context.JSON(200, gin.H{"users": users})
}