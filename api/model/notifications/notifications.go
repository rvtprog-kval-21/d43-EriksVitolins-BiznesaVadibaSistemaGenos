package notifications

import (
	"api/database"
	"api/model/user"
	"fmt"
	"time"
)

type Notifications struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	Topic     string    `json:"topic"`
	Owner     user.User `json:"owner" gorm:"foreignKey:OwnerID"`
	OwnerID   int       `json:"owner_id" gorm:"index"`
	Author    user.User `json:"author" gorm:"foreignKey:AuthorID"`
	AuthorID  int       `json:"author_id" gorm:"index"`
	Seen      bool      `json:"seen" gorm:"index"`
	Published time.Time `json:"published" gorm:"index"`
	Content   string    `json:"content"`
}

func SaveNotifications(notification Notifications) interface{} {
	results := database.DBConn.Create(&notification)
	return results.Error
}

func SeeNotifications(id interface{}) []Notifications {
	var notification []Notifications
	response := database.DBConn.Preload("Author").Preload("Owner").Where("owner_id = ?", id).Where("seen = ?", false).Order("id desc").Find(&notification)
	if response != nil {
		fmt.Println(response)
	}
	return notification
}


func SeeNotificationsCount(userId interface{}) int64 {
	var count int64
	database.DBConn.Model(&Notifications{}).Where("owner_id = ?", userId).Where("seen = ?", false).Count(&count)
	return count
}

func UpdateIsSeen(id interface{}, userId interface{}, isSeen bool) Notifications {
	var notification Notifications
	database.DBConn.Model(&Notifications{}).Where("owner_id = ?", userId).Where("id = ?", id).Update("seen", isSeen)
	return notification
}
