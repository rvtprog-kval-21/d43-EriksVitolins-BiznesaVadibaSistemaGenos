package user_controller

import (
	"golang-api/database"
	"golang-api/model"
)

// FindByEmail find user by their email address
func FindByEmail(email string) (*model.User, error) {
	var user model.User
	response := database.DBConn.Where("email = ?", email).First(&user)
	return &user, response.Error
}
