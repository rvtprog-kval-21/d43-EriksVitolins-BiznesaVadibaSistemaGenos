package timetable

import (
	"api/model/timetable"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)
type RequesSave struct {
	Dates []timetable.TimetableResponse `json:"dates"`
}
func Save(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	var request RequesSave
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	id := int(claims["id"].(float64))

	schedule := make([]timetable.Timetable, len(request.Dates))
	for i, iter := range request.Dates {
		schedule[i].Status = iter.Status
		schedule[i].UserID = id
		schedule[i].Date = iter.Date
		if iter.Status == "1" {
			schedule[i].StartTime = iter.StartTime.HH + ":" + iter.StartTime.MM
			schedule[i].EndTime = iter.EndTime.HH + ":" + iter.EndTime.MM
		}
	}

	timetable.SaveTimetable(schedule)
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

type RequestGet struct {
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}
func Get(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	var request RequestGet
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	arr := timetable.GetTimetable(claims["id"],request.StartDate, request.EndDate)
	schedule := make([]timetable.TimetableResponse, len(arr))
	for i, iter := range arr {
		schedule[i].ID = iter.ID
		schedule[i].Status = iter.Status
		schedule[i].UserID = iter.UserID
		schedule[i].Date = iter.Date
		if iter.Status == "1" {
			timeSplit := strings.Split(iter.StartTime, ":")
			schedule[i].StartTime.HH = timeSplit[0]
			schedule[i].StartTime.MM = timeSplit[1]
			timeSplit = strings.Split(iter.EndTime, ":")
			schedule[i].EndTime.HH = timeSplit[0]
			schedule[i].EndTime.MM = timeSplit[1]
		}
	}
	context.JSON(http.StatusOK, gin.H{"dates": schedule})
}

func Update(context *gin.Context) {
	claims := jwtParser.GetClaims(context)
	if claims == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error unparsing the token"})
		return
	}
	var request RequesSave
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't unmarshal json"})
		return
	}
	id := int(claims["id"].(float64))
	for _, iter := range request.Dates {
		var schedule timetable.Timetable
		schedule.ID = iter.ID
		schedule.Status = iter.Status
		if iter.UserID != id {
			context.JSON(403, gin.H{"error": "User ids didn't match"})
			return
		}
		schedule.UserID = id
		schedule.Date = iter.Date
		if iter.Status == "1" {
			schedule.StartTime = iter.StartTime.HH + ":" + iter.StartTime.MM
			schedule.EndTime = iter.EndTime.HH + ":" + iter.EndTime.MM
		}
		timetable.UpdateTimetable(schedule)
	}

	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

