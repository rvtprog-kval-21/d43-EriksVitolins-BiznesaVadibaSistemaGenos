package timetable

import (
	"api/database"
	"api/model/user"
	"time"
)

type Timetable struct {
	ID        int    `json:"id" gorm:"primaryKey;not null"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID    int       `json:"user_id" gorm:"index"`
	Date	time.Time `json:"date" gorm:"index"`
	Status string `json:"status"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
}

type TimetableResponse struct {
	ID        int    `json:"id" gorm:"primaryKey;not null"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID    int       `json:"user_id" gorm:"index"`
	Date	time.Time `json:"date" gorm:"index"`
	Status string `json:"status"`
	StartTime timeObj `json:"start_time"`
	EndTime timeObj `json:"end_time"`
}

type timeObj struct {
	HH string `json:"hh"`
	MM string `json:"mm"`
}

func SaveTimetable(tables []Timetable) {
	database.DBConn.Create(&tables)
}

func UpdateTimetable(table Timetable) {
	database.DBConn.Save(&table)
}

func GetTimetable(userID interface{},startDate time.Time, endDate time.Time) []TimetableResponse {
	var dates []Timetable
	database.DBConn.Where("user_id = ?", userID).Where("date >= ?", startDate).Where("date <= ?", endDate).Find(&dates)
	return []TimetableResponse{}
}