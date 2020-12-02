package projects

import (
	"api/database"
	"api/model/user"
	"gorm.io/gorm"
	"time"
)

type Project struct {
	ID        int            `gorm:"primaryKey;not null" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"index;not null" json:"name"`
	Avatar    string         `json:"avatar"`
	About     string         `json:"about"`
	Members   []Member       `json:"members"`
}

type Member struct {
	ProjectID int       `json:"project_id" gorm:"index"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID    int       `json:"user_id" gorm:"index"`
	IsAdmin   bool      `json:"is_admin" gorm:"default:0;index"`
	IsOwner   bool      `json:"is_owner" gorm:"default:0;index"`
}

func AddProject(project *Project) interface{} {
	results := database.DBConn.Create(&project)
	return results.Error
}

func UpdateAvatar(article *Project) interface{} {
	results := database.DBConn.Model(&article).Update("avatar", article.Avatar)
	return results.Error
}

func AddMembers(member *Member) interface{} {
	results := database.DBConn.Create(&member)
	return results.Error
}

func GetAll() []Project{
	var projects []Project
	database.DBConn.Preload("Members.User").Find(&projects)
	return projects
}