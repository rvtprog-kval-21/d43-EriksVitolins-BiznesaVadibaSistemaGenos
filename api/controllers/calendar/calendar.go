package calendar

import (
	"api/model/calendar"
	"api/model/notifications"
	"api/model/projects"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		} else {
			notificationNew := notifications.Notifications{
				AuthorID: int(claims["id"].(float64)),
				OwnerID: iter,
				Topic: "Calendar",
				Published: time.Now(),
				Seen: false,
				Content: "An event has been added to your calendar on " +
					event.StartDate.Format("2006-01-02 15:04:05") + " - " +
					event.EndDate.Format("2006-01-02 15:04:05") +
					" ,event is named " + event.Title,
			}
			notifications.SaveNotifications(notificationNew)
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

func GetEventsHome(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}

	events := calendar.GetEventsHome(time.Now(), time.Now().AddDate(0, 0, 7 * 1), claims["id"])

	context.JSON(200, gin.H{"events": events})
}

func DeleteEvent(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	memeber := calendar.GetMember(claims["id"],context.Param("id"))

	if memeber.ID == 0 || memeber.IsOwner == false {
		context.JSON(500, gin.H{"error": "You don't have access"})
	}

	calendar.DeleteAllEventsUsers(context.Param("id"))
	calendar.DeleteEvent(context.Param("id"))
	context.JSON(200, gin.H{"message": "success"})
}

func LeaveEvent(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	calendar.DeleteEventMember(context.Param("id"),claims["id"])
	context.JSON(200, gin.H{"message": "success"})
}


func UpdateEvent(context *gin.Context) {
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
	memeber := calendar.GetMember(claims["id"],context.Param("id"))

	if memeber.ID == 0 || memeber.IsOwner == false {
		context.JSON(500, gin.H{"error": "You don't have access"})
	}

	var ids []int
	ids = append(ids, int(claims["id"].(float64)))
	for _, iter := range event.Users {
		if arrayContains(ids, iter.ID) {
			continue
		}
		ids = append(ids, iter.ID)
	}

	event.StartDate = event.Time[0]
	event.EndDate = event.Time[1]

	event.ID, _ = strconv.Atoi(context.Param("id"))
	calendar.UpdateEvent(event)
	for _, iter := range ids {
		memeberOfEvent := calendar.GetMember(iter,context.Param("id"))
		if memeberOfEvent.ID != 0{
			continue
		}

		notificationNew := notifications.Notifications{
			AuthorID: int(claims["id"].(float64)),
			OwnerID: iter,
			Published: time.Now(),
			Seen: false,
			Content: "An event has been added to your calendar on " +
				event.StartDate.Format("2006-01-02 15:04:05") + " - " +
				event.EndDate.Format("2006-01-02 15:04:05") +
				" ,event is named " + event.Title,
		}
		notifications.SaveNotifications(notificationNew)
		memberNew := calendar.EventMember{
			IsOwner: false,
			EventID: event.ID,
			UserID:  iter,
		}
		calendar.SaveEventMember(memberNew)
	}
	calendar.DeleteEveryOtherMember(event.ID,ids)

	context.JSON(200, gin.H{"message": "success"})
}


func arrayContains(arr []int, needle int) bool {
	for _, iter := range arr {
		if iter == needle {
			return true
		}
	}
	return false
}
