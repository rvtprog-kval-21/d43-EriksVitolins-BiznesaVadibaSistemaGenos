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
	IsInitiated bool `json:"is_initiated" gorm:"default:0"`
}

type AnnouncementsUser struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	UserID    int       `json:"user_id" gorm:"foreignKey:UserID;index"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at" json:"created_at"`
}

type FollowingUser struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	Following      User      `json:"following" gorm:"foreignKey:FollowingID"`
	FollowingID    int       `json:"following_id" gorm:"foreignKey:FollowingID;index"`
	Follower      User      `json:"follower" gorm:"foreignKey:FollowerID"`
	FollowerID    int       `json:"follower_id" gorm:"foreignKey:FollowerID;index"`
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

func GetUserById(id interface{}) (*User, error) {
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

func UpdateAvatar(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Update("avatar", user.Avatar)
}

func UpdatePassword(user User) {
	database.DBConn.Model(&User{}).Where("email = ?", user.Email).Update("password", user.Password)
}

func UpdateNameAndLastName(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Updates(User{Name: user.Name, LastName: user.LastName})
}

func UpdateBirthdayAndNameday(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Updates(User{Birthday: user.Birthday, NameDay: user.NameDay})
}

func UpdateTitle(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Update("title", user.Title)
}

func UpdateNumber(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Update("phone_number", user.PhoneNumber)
}

func UpdateAbout(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Update("about", user.About)
}

func UpdateBackground(user User) {
	database.DBConn.Model(&User{}).Where("id = ?", user.ID).Update("background", user.Background)
}

func CreateUsers(user User) (interface{}, User) {
	response := database.DBConn.Create(&user)
	return response.Error, user
}

func SearchUsers(search string) []User {
	var users []User
	database.DBConn.Select("email", "id").Where("email like ?", "%"+search+"%").Find(&users)
	return users
}

func SaveAnnouncement(annc AnnouncementsUser) {
	database.DBConn.Create(&annc)
}

func GetAnnouncementUser(id string) []AnnouncementsUser {
	var annc []AnnouncementsUser
	database.DBConn.Where("user_id = ?", id).Order("id desc").Find(&annc)
	return annc
}

func GetFollowedAnnouncementUser(id interface{}) []AnnouncementsUser {
	var followers []FollowingUser
	var ids []int
	var annc []AnnouncementsUser
	database.DBConn.Where("follower_id = ?", id).Find(&followers)
	for _, iter := range followers{
		ids = append(ids, iter.FollowingID)
	}
	database.DBConn.Preload("User").Where("user_id in (?)", ids).Find(&annc).Order("id desc").Limit(10)
	return annc
}

func DeleteAnnouncementUser(id string, user_id interface{}) {
	database.DBConn.Where("user_id = ?", user_id).Where("id = ?", id).Delete(AnnouncementsUser{})
}

func DeleteFollowerUser(user FollowingUser) {
	database.DBConn.Where("follower_id = ?", user.FollowerID).Where("following_id = ?", user.FollowingID).Delete(FollowingUser{})
}

func AddFollowerUser(user FollowingUser) {
	database.DBConn.Create(&user)
}

func SearchForFollowerUser(user FollowingUser)  FollowingUser {
	database.DBConn.Where("follower_id = ?", user.FollowerID).Where("following_id = ?", user.FollowingID).Find(&user)
	return user
}

func SearchForUser(search string) []User{
	var users []User
	if search == "" {
		database.DBConn.Limit(10).Find(&users)
	} else {
		database.DBConn.Where("name LIKE ? OR last_name LIKE ? OR email LIKE ?", "%"+search +"%", "%"+search +"%", "%"+search +"%").Find(&users)
	}
	return users
}

func InitUser(id interface{})  {
	database.DBConn.Model(&User{}).Where("id = ?", id).Update("is_initiated",true)
}