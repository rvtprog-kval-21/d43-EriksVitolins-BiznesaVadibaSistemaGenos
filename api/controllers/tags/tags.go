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

type TagRequest struct {
	Name string `json:"name"`
	About string `json:"about"`
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

func JoinTag(context *gin.Context) {
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tags.JoinTag(context.Param("id"), claims["id"])
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	database.Close()
	context.JSON(201, gin.H{"message": "Congrats you joined this tag"})
}

func DeleteTag(context *gin.Context) {
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tags.DeleteTag(context.Param("id"), claims["id"])
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	database.Close()
	context.JSON(201, gin.H{"message": "Congrats you deleted this tag"})
}

func NewName(context *gin.Context) {
	var request TagRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "There was an error unparsing the json"})
		return
	}
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tags.UpdateName(context.Param("id"), claims["id"], request.Name )
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	database.Close()
	context.JSON(201, gin.H{"message": "You changed the name"})
}

func NewAbout(context *gin.Context) {
	var request TagRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "There was an error unparsing the json"})
		return
	}
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tags.UpdateAbout(context.Param("id"), claims["id"], request.About )
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	database.Close()
	context.JSON(201, gin.H{"message": "You changed the About"})
}

func MakePublic(context *gin.Context)  {
	var request TagRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "There was an error unparsing the json"})
		return
	}
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tags.UpdatePublic(context.Param("id"), claims["id"], true )
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	database.Close()
	context.JSON(201, gin.H{"message": "You changed the public setting"})
}

func MakePrivate(context *gin.Context)  {
	database.Open()
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tags.UpdatePublic(context.Param("id"), claims["id"], false )
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	database.Close()
	context.JSON(201, gin.H{"message": "You changed the public setting"})
}