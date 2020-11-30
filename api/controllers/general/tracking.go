package general

import (
	"api/model/tracking"
	"api/model/user"
	"api/utlis/jwtParser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
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
	newSubmission.Subject = context.PostForm("subject")
	newSubmission.Description = context.PostForm("description")
	newSubmission.SubmitDate = time.Now()
	response := tracking.AddSubmission(&newSubmission)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the blog"})
		return
	}
	context.JSON(201, gin.H{"message": "Submission is submitted", "id": newSubmission.ID})
}

func SeePersonalSubmissions(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	submissions := tracking.GetYourSubmissions(claims["id"])

	context.JSON(200, gin.H{"submissions": submissions})
}

func SeeSubmissions(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	manager := tracking.Profile(claims["id"])
	if manager.ID == 0 {
		context.JSON(403, gin.H{"error": "You don't have access"})
		return
	}
	submissions := tracking.GetSubmissions()

	context.JSON(200, gin.H{"submissions": submissions})
}

func AddAttachments(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	file, _ := context.FormFile("file")
	if file != nil {
		path := "/tracking/%s/"
		path = fmt.Sprintf(path, context.PostForm("id"))
		if _, err := os.Stat("storage" + path); os.IsNotExist(err) {
			os.MkdirAll("storage"+path, os.ModeDir)
		}
		path = path + file.Filename
		err := context.SaveUploadedFile(file, "storage"+path)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the banner"})
			return
		}
		var newAttachment tracking.TrackedAttachment
		newAttachment.Path = path
		newAttachment.SubmissionID, _ = strconv.Atoi(context.PostForm("id"))
		response := tracking.AddTracking(&newAttachment)
		if response != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error saving the blog banner"})
			return
		}
	}
	context.JSON(201, gin.H{"message": "Article is saved"})
}

func OpenSubmissions(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	manager := tracking.Profile(claims["id"])
	if manager.ID == 0 {
		context.JSON(403, gin.H{"error": "You don't have access"})
		return
	}
	err := tracking.OpenSubmissions(context.Param("id"))
	if err != nil {
		context.JSON(500, gin.H{"error": "Had trouble updating submission"})
	}
	context.JSON(200, "good")
}

func CloseSubmissions(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	manager := tracking.Profile(claims["id"])
	if manager.ID == 0 {
		context.JSON(403, gin.H{"error": "You don't have access"})
		return
	}
	err := tracking.CloseSubmissions(context.Param("id"))
	if err != nil {
		context.JSON(500, gin.H{"error": "Had trouble updating submission"})
	}
	context.JSON(200, "good")
}
