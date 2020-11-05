package tags

import (
	"api/database"
	"api/model/tags"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseIndex struct {
	Tags *[]tags.Tag `json:"tags"`
}

func IndexPublic(context *gin.Context) {
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	tagsObject := tags.GetAllPublicTags()
	database.Close()
	context.JSON(http.StatusOK, ResponseIndex{Tags: &tagsObject})
}

func IndexPrivate(context *gin.Context) {
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	tagsObject := tags.GetAllMemberTags(claims["id"])
	database.Close()
	context.JSON(http.StatusOK,  ResponseIndex{Tags: &tagsObject})
}
