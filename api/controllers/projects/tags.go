package projects

import (
	"api/model/projects"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type tagRequest struct {
	Tags []string `json:"tags"`
}

func ChangeTags(context *gin.Context) {
	var tags tagRequest
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user, errs := projects.GetMember(context.Param("id"), claims["id"])
	if errs != nil && !(user.IsOwner || user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	err := context.ShouldBindJSON(&tags)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	for _, iter := range tags.Tags {
		var newTag projects.Tag
		newTag.Name = iter
		newTag.ProjectID, _ = strconv.Atoi(context.Param("id"))
		tag, _ := projects.GetTag(newTag.Name, context.Param("id"))
		if tag.ID != 0 {
			continue
		}
		response := projects.AddTag(&newTag)
		if response != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't save all tags"})
			return
		}
	}
	projects.RemoveOrphans(tags.Tags, context.Param("id"))
	context.JSON(200, gin.H{"message": "Members Added"})
}

func UpdateTags(context *gin.Context) {
	var tag projects.Tag
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user, errs := projects.GetMember(context.Param("id"), claims["id"])
	if errs != nil && !(user.IsOwner || user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	err := context.ShouldBindJSON(&tag)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	response := projects.UpdateMemberTags(context.Param("id"), context.Query("id"), tag.ID)
	if response != nil {
		context.JSON(500, gin.H{"message": "There was an error"})
		return
	}
	context.JSON(200, gin.H{"message": "Tag changed"})
}
