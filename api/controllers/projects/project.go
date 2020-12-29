package projects

import (
	"api/model/projects"
	"api/model/user"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Request struct {
	Name  string `json:"name"`
	About string `json:"about"`
}

type RequestInvitees struct {
	Users []user.User `json:"users"`
}

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

func ArchiveProject(context *gin.Context) {
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
	err = projects.DeleteProject(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Member was deleted"})
}

func ChangeName(context *gin.Context) {
	var request Request
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
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
	if request.Name == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Field is empty"})
		return
	}
	var newProject projects.Project
	integer, err := strconv.Atoi(context.Param("id"))
	newProject.ID = integer
	newProject.Name = request.Name
	response := projects.UpdateName(&newProject)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	context.JSON(200, gin.H{"message": "Name was changed successfully"})
}

func ChangeAbout(context *gin.Context) {
	var request Request
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
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
		if request.About == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Field is empty"})
		return
	}
	var newProject projects.Project

	integer, err := strconv.Atoi(context.Param("id"))
	newProject.ID = integer
	newProject.About = request.About
	response := projects.UpdateAbout(&newProject)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}

	context.JSON(200, gin.H{"message": "Name was changed successfully"})
}

func ChangeAvatar(context *gin.Context) {
	var newProject projects.Project
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
	file, _ := context.FormFile("avatar")
	if file != nil {
		integer, err := strconv.Atoi(context.Param("id"))
		newProject.ID = integer
		path := "/projects/%s/"
		path = fmt.Sprintf(path, fmt.Sprint(newProject.ID))
		if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
			os.MkdirAll("storage"+path, os.ModeDir)
		}
		path = path + file.Filename
		err = context.SaveUploadedFile(file, "storage"+path)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the avatar"})
			return
		}
		newProject.Avatar = path
		response := projects.UpdateAvatar(&newProject)
		if response != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the Avatar"})
			return
		}
		context.JSON(200, gin.H{"message": "Avatar was changed successfully"})
	}
	context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error getting the Avatar"})
	return
}

func GetNonMembers(context *gin.Context) {
	user := projects.GetNonMembers(context.Param("id"))
	context.JSON(200,user)
}

func AddUsers(context *gin.Context) {
	var users RequestInvitees
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
	err := context.ShouldBindJSON(&users)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	for _, iter := range users.Users {
		var newMember projects.Member
		newMember.UserID = iter.ID
		newMember.IsOwner = false
		newMember.IsAdmin = false
		newMember.ProjectID, _ = strconv.Atoi(context.Param("id"))
		response := projects.AddMembers(&newMember)
		if response != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the Members"})
			return
		}
	}
	context.JSON(200, gin.H{"message": "Members Added"})
}