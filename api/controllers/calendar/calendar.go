package calendar

import (
	"api/model/calendar"
	"api/model/projects"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EventsRequest struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func CreateEvent(context *gin.Context) {
	var event calendar.Event
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	var idsTags []int
	for _, iter := range event.Tags {
		if arrayContains(idsTags, iter.ID) {
			continue
		}
		idsTags = append(idsTags, iter.ID)
	}
	MemebersOfTags := projects.GetMembersOfTags(idsTags)

	var ids []int
	ids = append(ids, int(claims["id"].(float64)))
	for _, iter := range event.Users {
		if arrayContains(ids, iter.ID) {
			continue
		}
		ids = append(ids, iter.ID)
	}

	for _, iter := range MemebersOfTags {
		if arrayContains(ids, iter.UserID) {
			continue
		}
		ids = append(ids, iter.UserID)
	}

	event.StartDate = event.Time[0]
	event.EndDate = event.Time[1]

	_, id := calendar.SaveEvent(event)
	for _, iter := range ids {
		isOwner := false
		if iter == int(claims["id"].(float64)) {
			isOwner = true
		}
		member := calendar.EventMember{
			IsOwner: isOwner,
			EventID: id,
			UserID:  iter,
		}
		calendar.SaveEventMember(member)
	}

	context.JSON(200, gin.H{"message": "success"})
}

func GetEvents(context *gin.Context) {
	var eventData EventsRequest
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	err := context.ShouldBindJSON(&eventData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}

	events := calendar.GetEvents(eventData.StartDate, eventData.EndDate, claims["id"])

	context.JSON(200, gin.H{"events": events})
}

func arrayContains(arr []int, needle int) bool {
	for _, iter := range arr {
		if iter == needle {
			return true
		}
	}
	return false
}
