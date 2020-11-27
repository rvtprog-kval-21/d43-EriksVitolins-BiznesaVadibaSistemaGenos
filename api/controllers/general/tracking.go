package general

import (
	"api/model/tracking"
	"api/model/user"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func GetAllManagers(context *gin.Context) {
	managers := tracking.AllManagers()
	var users []user.User
	for _, iter := range managers {
		users = append(users, iter.User)
	}

	context.JSON(200, gin.H{"users": users})
}

func AddUserToManagerRole(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tracking.JoinRole(context.Param("id"))
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "Congrats user is now a manager"})
}

func DeleteUserFromManagerRole(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := tracking.DeleteRole(context.Param("id"))
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": response})
		return
	}
	context.JSON(201, gin.H{"message": "User isn't a manager anymore"})
}

func ManagerIsMemeber(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	manager := tracking.Profile(claims["id"])
	if manager.ID == 0 {
		context.JSON(200, gin.H{"isManager": false})
		return
	}
	context.JSON(200, gin.H{"isManager": true})
}

func AddSubmission(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	var newSubmission tracking.TrackedSubmission
	newSubmission.UserID = int(claims["id"].(float64))
	newSubmission.Subject = context.PostForm("title")
	newSubmission.Description = context.PostForm("topic")
	newSubmission.SubmitDate = time.Now()
	newSubmission.IsConfirmed = false
	response := tracking.AddSubmission(&newSubmission)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the blog"})
		return
	}
	form, _ := context.MultipartForm()
	files, _ := form.File["files[]"]
	path := "/blog/%s"
	path = fmt.Sprintf(path, fmt.Sprint(newSubmission.ID))
	if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
		os.MkdirAll("storage"+path, os.ModeDir)
	}
	path = path + "/banner.png"
	err := context.SaveUploadedFile(files[0], path)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the banner"})
		return
	}
	context.JSON(201, gin.H{"message": "Submission is submitted"})
}

func SeePersonalSubmissions(context *gin.Context) {

}
