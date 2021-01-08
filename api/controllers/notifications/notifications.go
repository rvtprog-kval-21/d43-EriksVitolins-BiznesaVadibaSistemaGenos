package notifications

import (
	"api/model/notifications"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotifsAvailable(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}

	count := notifications.SeeNotificationsCount(claims["id"])

	context.JSON(http.StatusOK, gin.H{"count": count})
}

func Notifs(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	response := notifications.SeeNotifications(claims["id"])

	context.JSON(http.StatusOK, gin.H{"notifications": response})
}

func UpdateSeenToTrue(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	notifications.UpdateIsSeen(context.Param("id"),claims["id"],true)

	context.JSON(http.StatusOK, gin.H{"Message": "All good"})
}