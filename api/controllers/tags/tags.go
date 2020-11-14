package tags

import (
	"api/database"
	"api/model/tags"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseIndex struct {
	Tags *[]tags.Tag `json:"tags"`
}

type TagRequest struct {
	Name  string `json:"name"`
	About string `json:"about"`
}
type AvatarRequest struct {
	File string `json:"file"`
}

func IndexPublic(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	tagsObject := tags.GetAllPublicTags()
	database.Close()
	context.JSON(http.StatusOK, ResponseIndex{Tags: &tagsObject})
}

func IndexPrivate(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	tagsObject := tags.GetAllMemberTags(claims["id"], false)
	database.Close()
	context.JSON(http.StatusOK, ResponseIndex{Tags: &tagsObject})
}

func Profile(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	tag, isMember := tags.Profile(context.Param("id"), claims["id"])
	database.Close()
	if tag.ID == 0 {
		context.JSON(500, gin.H{"error": "Coudn't find the record"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"tag": &tag, "is_member": isMember})
}

func JoinTag(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.JoinTag(context.Param("id"), claims["id"])
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "Congrats you joined this tag"})
}

func DeleteTag(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.DeleteTag(context.Param("id"), claims["id"])
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "Congrats you deleted this tag"})
}

func NewName(context *gin.Context) {
	var request TagRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "There was an error unparsing the json"})
		return
	}
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.UpdateName(context.Param("id"), claims["id"], request.Name)
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "You changed the name"})
}

func NewAbout(context *gin.Context) {
	var request TagRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "There was an error unparsing the json"})
		return
	}
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.UpdateAbout(context.Param("id"), claims["id"], request.About)
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "You changed the About"})
}

func MakePublic(context *gin.Context) {
	var request TagRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "There was an error unparsing the json"})
		return
	}
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.UpdatePublic(context.Param("id"), claims["id"], true)
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "You changed the public setting"})
}

func MakePrivate(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.UpdatePublic(context.Param("id"), claims["id"], false)
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "You changed the public setting"})
}

func LeaveTag(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	response := tags.LeaveTag(context.Param("id"), claims["id"])
	database.Close()
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "Congrats you left this tag"})
}

func SetAvatar(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	database.Open()
	isUser, isAdmin := tags.IsAdmin(context.Param("id"), claims["id"])
	database.Close()
	if !isUser || !isAdmin {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "You don't have access to make this change"})
		return
	}
	file, _ := context.FormFile("file")
	path := "./storage/group/%s/logo/logo.png"
	err := context.SaveUploadedFile(file, fmt.Sprintf(path, context.Param("id")))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the file"})
	}

	context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
