package tracking

import (
	"api/database"
	"api/model/user"
	"strconv"
)

type Managers struct {
	ID     int       `gorm:"primaryKey;not null" json:"id"`
	User   user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID int       `json:"user_id" gorm:"foreignKey:UserID;index"`
}

func AllManagers() []Managers {
	var managers []Managers
	database.DBConn.Preload("User").Find(&managers)
	return managers
}

func JoinRole(id string) interface{} {
	manager := Profile(id)
	if manager.ID != 0 {
		return "Already a manager"
	}
	intID, _ := strconv.Atoi(id)
	manager.UserID = intID
	results := database.DBConn.Create(&manager)
	return results.Error
}

func DeleteRole(id string) interface{} {
	manager := Profile(id)
	if manager.ID == 0 {
		return "Not a manager"
	}
	results := database.DBConn.Unscoped().Delete(&manager)
	return results.Error
}

func Profile(id interface{}) Managers {
	var manager Managers
	database.DBConn.Where("user_id = ?", id).Find(&manager)
	return manager
}