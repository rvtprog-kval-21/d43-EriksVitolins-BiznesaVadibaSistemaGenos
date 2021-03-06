package projects

import (
	"api/model/notifications"
	"api/model/projects"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type announcementRequest struct {
	Content string `json:"content"`
	Tags []projects.Tag `json:"tags"`
}

func SaveAnnouncement(context *gin.Context) {
	var request announcementRequest
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
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}

	var newAnnounc projects.ProjectAnnouncements
	newAnnounc.ProjectID, _ = strconv.Atoi(context.Param("id"))
	newAnnounc.Content = request.Content
	newAnnounc.AuthorID = int(claims["id"].(float64))
	newAnnounc.Published = time.Now()

	var tags []string
	for _, iter := range request.Tags{
		tags = append(tags, iter.Name)
	}
	ids := projects.GetTagMembers( context.Param("id"),tags)
	for _, iter := range ids {
		newNotification := notifications.Notifications{
			Published: time.Now(),
			Topic: "Announcement",
			OwnerID: iter.ID,
			AuthorID: int(claims["id"].(float64)),
			Content: request.Content,
		}
		notifications.SaveNotifications(newNotification)
	}
	response := projects.SaveAnnouncement(newAnnounc)
	if response != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}



	context.JSON(http.StatusOK, gin.H{"message": "Announcement was saved"})
}

func SeeAnnouncements(context *gin.Context) {
	current, _ := strconv.Atoi(context.Query("currentPage"))
	next := current + 1
	response := projects.SeeAnnouncement(current, next, context.Param("id"))

	context.JSON(http.StatusOK, gin.H{"annc": response})
}
