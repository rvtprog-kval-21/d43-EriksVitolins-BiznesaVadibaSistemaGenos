package timetable

import (
	"api/model/timetable"
	"api/utlis/jwtParser"
	"github.com/gin-gonic/gin"
	"net/http"
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
	context.JSON(http.StatusOK, gin.H{"dates": arr})
}

