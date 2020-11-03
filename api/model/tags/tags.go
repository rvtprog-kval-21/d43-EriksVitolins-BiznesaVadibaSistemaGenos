package tags

import (
	"api/model/user"
	"gorm.io/gorm"
	"time"
)

type Tags struct {
	ID        uint           `gorm:"primaryKey;not null" json:"id"`
	CreatedAt time.Time      `json:"created_at" json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserOwner user.User      `gorm:"foreignKey:Owner"`
	Owner     int            `json:"owner"`
	Name      string         `gorm:"index;not null" json:"id"`
	Avatar    string         `json:"avatar"`
	About     string         `json:"about"`
	Private   bool           `json:"private"`
	IsPublic  bool           `json:"is_public"`
}
