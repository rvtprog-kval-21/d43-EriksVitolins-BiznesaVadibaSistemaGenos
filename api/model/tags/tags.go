package tags

import (
	"api/database"
	"api/model/user"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	ID        int            `gorm:"primaryKey;not null" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `gorm:"index;not null" json:"name"`
	Avatar    string         `json:"avatar"`
	About     string         `json:"about"`
	IsPublic  bool           `json:"is_public" gorm:"default:1"`
	Members   []Member       `json:"members"`
}

type Member struct {
	TagID   int       `json:"tags_id" gorm:"index"`
	User    user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID  int
	IsAdmin bool `json:"is_admin" gorm:"default:0"`
	IsOwner bool `json:"is_owner" gorm:"default:0"`
}

func GetAllPublicTags() []Tag {
	var tags []Tag
	database.DBConn.Preload("Members.User").Where("is_public = ?", 1).Find(&tags)
	return tags
}

func GetAllMemberTags(id interface{}) []Tag {
	var members []Member
	var tags []Tag
	database.DBConn.Where("user_id = ?", id).Find(&members)
	data := NameList(members)
	database.DBConn.Preload("Members.User").Where("id in ?", data).Find(&tags)
	return tags
}

func NameList(u []Member) []int {
	var list []int
	for _, iter := range u {
		list = append(list, iter.TagID)
	}
	return list
}
