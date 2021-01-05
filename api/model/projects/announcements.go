package projects

import (
	"api/database"
	"api/model/user"
	"time"
)

type ProjectAnnouncements struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	ProjectID int       `json:"project_id" gorm:"index"`
	Author    user.User `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID  int       `json:"author_id" gorm:"index"`
	Published time.Time `json:"published" gorm:"index"`
	Content   string    `json:"content"`
}

func SaveAnnouncement(ann ProjectAnnouncements) interface{} {
	results := database.DBConn.Create(&ann)
	return results.Error
}

func SeeAnnouncement(currentPage int, NextPage int, id string) []ProjectAnnouncements {
	var annc []ProjectAnnouncements
	database.DBConn.Preload("Author").Where("project_id = ?", id).Offset(currentPage * 5).Order("id desc").Limit(NextPage * 3).Find(&annc)
	return annc
}
