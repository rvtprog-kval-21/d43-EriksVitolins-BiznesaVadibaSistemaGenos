package calendar

import (
	"api/database"
	"api/model/projects"
	"api/model/user"
	"time"
)

type Calendar struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	StartDate time.Time `json:"startDate" gorm:"index"`
	EndDate time.Time `json:"endDate" gorm:"index"`
	Title string `json:"title"`
	Description string `json:"description"`
	Public bool `json:"public" gorm:"index"`
	Owner     user.User `json:"owner" gorm:"foreignKey:OwnerID"`
	OwnerID   int       `json:"owner_id" gorm:"index"`
	Author    user.User `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID  int       `json:"author_id" gorm:"index"`
	Time []string `json:"time" gorm:"-"`
	Users []user.User `json:"users" gorm:"-"`
	Tags []projects.Tag `json:"tags" gorm:"-"`
}

func SaveEvent(event Calendar) interface{} {
	results := database.DBConn.Create(&event)
	return results.Error
}