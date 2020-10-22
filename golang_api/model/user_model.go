package user_model

import (
	"golang-api/database"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Avatar    string
	Name      string
	LastName  string
	About     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FindByEmail find user by their email address
func FindByEmail(email string) (*User, error) {
	var user User
	response := database.DBConn.Where("email = ?", email).First(&user)
	return &user, response.Error
}
