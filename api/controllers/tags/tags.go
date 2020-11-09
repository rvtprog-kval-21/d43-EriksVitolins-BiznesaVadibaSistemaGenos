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
	tagsObject := tags.GetAllMemberTags(claims["id"], false)
	database.Close()
	context.JSON(http.StatusOK, ResponseIndex{Tags: &tagsObject})
}

func Profile(context *gin.Context) {
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	tag, isMember := tags.Profile(context.Param("id"), claims["id"])
	database.Close()
	if tag.ID == 0 {
		context.JSON(500, gin.H{"error": "Coudn't find the record"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"tag": &tag, "is_member": isMember})
}
