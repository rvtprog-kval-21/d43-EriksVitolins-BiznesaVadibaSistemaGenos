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
	User    user.User `json:"user"`
	UserID  int
	IsAdmin bool `json:"is_admin" gorm:"default:0"`
	IsOwner bool `json:"is_owner" gorm:"default:0"`
}

func GetAllPublicTags() []Tag {
	var tags []Tag
	database.DBConn.Preload("Members").Where("is_public = ?", 1).Find(&tags)
	return tags
}

func GetAllMemberTags(id interface{}) []Tag {
	//	var tags []Tags
	//	database.DBConn.Where("is_public = ?", id).Find(&tags)
	//	return tags
	return nil
}
