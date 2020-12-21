package projects

import (
	"api/model/projects"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func CreateProject(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	var newProject projects.Project
	users := strings.Split(context.PostForm("users"), ";")
	newProject.Name = context.PostForm("name")
	newProject.About = context.PostForm("about")
	response := projects.AddProject(&newProject)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the project"})
		return
	}
	file, _ := context.FormFile("avatar")
	if file != nil {
		path := "/projects/%s/"
		path = fmt.Sprintf(path, fmt.Sprint(newProject.ID))
		if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
			os.MkdirAll("storage"+path, os.ModeDir)
		}
		path = path + file.Filename
		err := context.SaveUploadedFile(file, "storage"+path)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the avatar"})
			return
		}
		newProject.Avatar = path
		response = projects.UpdateAvatar(&newProject)
		if response != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the Avatar"})
			return
		}

	}
	for _, iter := range users {
		var newMember projects.Member
		newMember.UserID, _ = strconv.Atoi(iter)
		newMember.IsOwner = true
		newMember.IsAdmin = true
		newMember.ProjectID = newProject.ID
		response := projects.AddMembers(&newMember)
		if response != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the Members"})
			return
		}
	}
	context.JSON(201, gin.H{"message": "Project is created"})
}

func GetAll(context *gin.Context) {
	projectsArray := projects.GetAll()
	context.JSON(http.StatusOK, gin.H{"projects": projectsArray})
}

func GetProject(context *gin.Context) {
	project, err := projects.GetProject(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"project": "Doesn't exist"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"project": project})
}

func MakeAdmin(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user, err := projects.GetMember(context.Param("id"), claims["id"])
	if err != nil && !user.IsOwner {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	err = projects.UpdateAdmin(context.Param("id"), context.Query("id"), true)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Admin rights were granted"})
}

func UnMakeAdmin(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user, err := projects.GetMember(context.Param("id"), claims["id"])
	if err != nil && !user.IsOwner {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	err = projects.UpdateAdmin(context.Param("id"), context.Query("id"), false)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Admin rights were taken away"})
}

func KickUser(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	user, err := projects.GetMember(context.Param("id"), claims["id"])
	if err != nil && !(user.IsOwner || user.IsAdmin) {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	err = projects.DeleteMember(context.Param("id"), context.Query("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Member was deleted"})
}

func LeaveProject(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	_, err := projects.GetMember(context.Param("id"), claims["id"])
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	err = projects.DeleteMember(context.Param("id"), claims["id"])
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	project, err := projects.GetProject(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"project": "Doesn't exist"})
		return
	}
	if len(project.Members) > 0 {
		context.JSON(http.StatusOK, gin.H{"message": "Member left successfully"})
	} else {
		err = projects.DeleteProject(context.Param("id"))
		if err != nil {
			context.JSON(200, gin.H{"error": "Project coudn't be deleted"})
		} else {
			context.JSON(200, gin.H{"error": "Project was deleted"})
		}
	}
}
