package bloggers

import (
	"api/database"
	"api/model/user"
	"strconv"
)

type Bloggers struct {
	ID        int            `gorm:"primaryKey;not null" json:"id"`
	User    user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID  int `json:"user_id" gorm:"foreignKey:UserID;index"`
}

func JoinRole(id string) interface{} {
	blogger := Profile(id)
	if blogger.ID != 0 {
		return "Already a member"
	}
	intID, _ := strconv.Atoi(id)
	blogger.UserID = intID
	results := database.DBConn.Create(&blogger)
	return results.Error
}

func DeleteRole(id string) interface{} {
	blogger := Profile(id)
	if blogger.ID == 0 {
		return "Not a blogger"
	}
	results := database.DBConn.Unscoped().Delete(&blogger)
	return results.Error
}

func Profile(id string) Bloggers {
	var blogger Bloggers
	database.DBConn.Where("user_id = ?", id).Find(&blogger)

	return blogger
}

func All() []Bloggers {
	var blogger []Bloggers
	database.DBConn.Preload("User").Find(&blogger)
	return blogger
}