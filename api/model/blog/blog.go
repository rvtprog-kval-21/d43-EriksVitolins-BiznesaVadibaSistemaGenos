package blog

import (
	"api/database"
	"api/model/user"
	"gorm.io/gorm"
	"time"
)

type Blogs struct {
	ID        uint           `gorm:"primaryKey" json:"id" form:"id"`
	PublishAt time.Time      `json:"publish_at" gorm:"index" form:"publish_at"`
	User      user.User      `json:"user" gorm:"foreignKey:UserID"`
	UserID    int            `json:"user_id" gorm:"foreignKey:UserID;index"`
	Title     string         `json:"title" gorm:"index" form:"title"`
	Headtext  string         `json:"headtext" form:"headtext"`
	Photo     string         `json:"photo"`
	Topic     string         `json:"topic" form:"topic"`
	Content   string         `json:"content" form:"content"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func AddBlog(article *Blogs) interface{}  {
	results := database.DBConn.Create(&article)
	return results.Error
}

func UpdateBannerBlog(article *Blogs) interface{}  {
	results := database.DBConn.Model(&article).Update("photo", article.Photo)
	return results.Error
}