package blog

import (
	"api/model/user"
	"gorm.io/gorm"
	"time"
)

type Blogs struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PublishAt time.Time      `json:"publish_at" gorm:"index"`
	User      user.User      `json:"user" gorm:"foreignKey:UserID"`
	UserID    int            `json:"user_id" gorm:"foreignKey:UserID;index"`
	Title     string         `json:"title" gorm:"index"`
	Headtext  string         `json:"headtext"`
	Photo     string         `json:"photo"`
	Topic     string         `json:"topic"`
	Content   string         `json:"content"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
