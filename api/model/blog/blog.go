package blog

import (
	"api/database"
	"api/model/user"
	"gorm.io/gorm"
	"time"
)

type Blogs struct {
	ID        int           `gorm:"primaryKey" json:"id" form:"id"`
	PublishAt time.Time      `json:"publish_at" gorm:"index" form:"publish_at"`
	User      user.User      `json:"user" gorm:"foreignKey:UserID"`
	UserID    int            `json:"user_id" gorm:"foreignKey:UserID;index"`
	Title     string         `json:"title" gorm:"index" form:"title"`
	Headtext  string         `json:"headtext" form:"headtext"`
	Photo     string         `json:"photo"`
	Topic     string         `json:"topic" form:"topic"`
	Content   string         `json:"content" form:"content"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func AddBlog(article *Blogs) interface{} {
	results := database.DBConn.Create(&article)
	return results.Error
}

func UpdateBannerBlog(article *Blogs) interface{} {
	results := database.DBConn.Model(&article).Update("photo", article.Photo)
	return results.Error
}

func GetYourBlogs(userID interface{}) ([]Blogs, interface{}) {
	var blogs []Blogs
	response := database.DBConn.Unscoped().Where("user_id = ?", userID).Find(&blogs)
	return blogs, response.Error
}

func GetBlogUnscoped(id interface{}) Blogs {
	var blog Blogs
	database.DBConn.Unscoped().Preload("User").Where("id = ?", id).First(&blog)
	return blog
}

func DeleteBlog(blog *Blogs) interface{} {
	results := database.DBConn.Delete(&blog)
	return results.Error
}

func UnDeleteBlog(id interface{}) interface{} {
	response := database.DBConn.Model(&Blogs{}).Where("id = ?", id).Update("deleted_at", nil)
	return response.Error
}

func UpdateBlog(blogs *Blogs) interface{} {
	results := database.DBConn.Model(&Blogs{}).Where("id = ?", blogs.ID).Where("user_id = ?", blogs.UserID).Updates(Blogs{Title: blogs.Title, Topic: blogs.Topic, Content: blogs.Content, Headtext: blogs.Headtext, PublishAt: blogs.PublishAt})
	return results.Error
}

func GetBlogs() ([]Blogs, interface{}) {
	var blogs []Blogs
	response := database.DBConn.Preload("User").Order("id desc").Where("DATE(publish_at) < DATE(?)", time.Now()).Find(&blogs)
	return blogs, response.Error
}

func GetBlogsLimitedToFour() ([]Blogs, interface{}) {
	var blogs []Blogs
	response := database.DBConn.Preload("User").Order("id desc").Where("DATE(publish_at) < DATE(?)", time.Now()).Limit(4).Find(&blogs)
	return blogs, response.Error
}
