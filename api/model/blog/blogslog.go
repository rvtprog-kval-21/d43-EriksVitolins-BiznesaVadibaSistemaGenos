package blog

import "api/database"

type BlogsLog struct {
	ID        int           `gorm:"primaryKey" json:"id"`
	UserID int       `json:"user_id" gorm:"index"`
	BlogID int       `json:"blog_id" gorm:"index"`

}

func AddItemToLog(userId interface{}, blogId interface{}) interface{} {
	response := database.DBConn.Create(&BlogsLog{
		UserID: userId.(int),
		BlogID: blogId.(int),
	})
	return response.Error
}

func GetCount(blogId interface{}) (int64, interface{})  {
	var count int64
	var blogs []BlogsLog
	response := database.DBConn.Where("blog_id = ?", blogId).Find(&blogs).Count(&count)
	return count, response.Error
}