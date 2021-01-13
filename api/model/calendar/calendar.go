package calendar

import (
	"api/database"
	"api/model/projects"
	"api/model/user"
	"time"
)

type Event struct {
	ID          int            `gorm:"primaryKey;not null" json:"id"`
	StartDate   time.Time      `json:"startDate" gorm:"index"`
	EndDate     time.Time      `json:"endDate" gorm:"index"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Public      bool           `json:"public" gorm:"index"`
	Time        []time.Time    `json:"time" gorm:"-"`
	Users       []user.User    `json:"users" gorm:"-"`
	Tags        []projects.Tag `json:"tags" gorm:"-"`
	Members     []EventMember  `json:"members"`
}

type EventMember struct {
	ID      int       `gorm:"primaryKey;not null" json:"id"`
	User    user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID  int       `json:"user_id" gorm:"index"`
	Event   Event     `json:"event" gorm:"foreignKey:EventID"`
	EventID int       `json:"event_id" gorm:"index"`
	IsOwner bool      `json:"isOwner"`
}

func SaveEvent(event Event) (interface{}, int) {
	results := database.DBConn.Create(&event)
	return results.Error, event.ID
}

func SaveEventMember(member EventMember) interface{} {
	results := database.DBConn.Create(&member)
	return results.Error
}

func GetEvents(startDate time.Time, endDate time.Time, id interface{}) []Event {
	var events []Event
	database.DBConn.Preload("Members.User").Joins("Left Join event_members ON event_members.event_id = events.id").
		Where("events.start_date > ? OR events.end_date < ?", startDate, endDate).
		Where("events.public = TRUE OR event_members.user_id = ?", id).
		Find(&events)

	return events
}
