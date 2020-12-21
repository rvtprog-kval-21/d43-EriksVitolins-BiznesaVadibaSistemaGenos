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

func GetAll() []Project {
	var projects []Project
	database.DBConn.Preload("Members.User").Find(&projects)
	return projects
}

func GetProject(id interface{}) (Project, interface{}) {
	var project Project
	response := database.DBConn.Preload("Members.User").Where("id = ?", id).Find(&project)
	return project, response.Error
}

func GetMember(id interface{}, userID interface{}) (Member, interface{}) {
	var member Member
	response := database.DBConn.Where("project_id = ?", id).Where("user_id = ?", userID).Find(&member)
	return member, response.Error
}

func UpdateAdmin(id interface{}, userID interface{}, isAdmin bool) interface{} {
	response := database.DBConn.Model(&Member{}).Where("project_id = ?", id).Where("user_id = ?", userID).Update("is_admin", isAdmin)
	return response.Error
}

func DeleteMember(id interface{}, userID interface{}) interface{} {
	response := database.DBConn.Where("project_id = ?", id).Where("user_id = ?", userID).Delete(Member{})
	return response.Error
}

func DeleteProject(id interface{}) interface{} {
	response := database.DBConn.Where("id = ?", id).Delete(Project{})
	return response.Error
}
