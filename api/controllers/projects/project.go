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
		}
	}
	context.JSON(201, gin.H{"message": "Project is created"})
}

func GetAll(context *gin.Context)  {
	projectsArray := projects.GetAll()
	context.JSON(http.StatusOK, gin.H{"projects": projectsArray})
}