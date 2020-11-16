package user

import (
	"api/database"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID          int            `gorm:"primaryKey;not null" json:"id"`
	CreatedAt   time.Time      `json:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Email       string         `gorm:"unique;not null" json:"email"`
	Password    string         `gorm:"not null" json:"-"`
	Role        string         `gorm:"not null" json:"role"`
	Avatar      string         `json:"avatar"`
	Background  string         `json:"background"`
	Name        string         `json:"name"`
	LastName    string         `json:"last_name"`
	About       string         `json:"about"`
	Title       string         `json:"title"`
	PhoneNumber string         `json:"phone_number"`
	Birthday    time.Time      `json:"birthday"`
	NameDay     time.Time      `json:"name_day"`
}

// FindByEmail find user by their email address
func FindByEmail(email string) (*User, error) {
	var user User
	response := database.DBConn.Select("email", "password", "id", "Role", "avatar", "name", "last_name").Where("email = ?", email).First(&user)
	return &user, response.Error
}

func GetAllUsers() []User {
	var users []User
	database.DBConn.Unscoped().Find(&users)
	return users
}

func GetUserById(id string) (*User, error) {
	var user User
	response := database.DBConn.Unscoped().Where("id = ?", id).First(&user)
	return &user, response.Error
}

func SoftDeleteUser(id string) error {
	response := database.DBConn.Where("id = ?", id).Delete(&User{})
	return response.Error
}

func UnlockUser(id string) error {
	response := database.DBConn.Model(&User{}).Where("id = ?", id).Update("deleted_at", nil)
	return response.Error
}

func NewEmail(id *string, email *string) (bool, string) {
	user, _ := FindByEmail(*email)
	if user.Email == "" {
		response := database.DBConn.Model(&User{}).Where("id = ?", *id).Update("email", *email)
		if response.Error != nil {
			return false, "Error with database"
		} else {
			return true, ""
		}
	} else {
		return false, "Email already taken"
	}
	return false, "Internal error"
}

func GetUsersIn(ids []interface{}) []User {
	var users []User
	database.DBConn.Select("email", "id", "avatar", "name", "last_name").Where("id in ?", ids).Find(&users)
	return users
}