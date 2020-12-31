package projects

import "api/model/user"

type Tags struct {
	ID int `json:"id" gorm:"primaryKey;not null"`
	Name string `json:"name"  gorm:"index;not null"`
	ProjectID    int       `json:"project_id" gorm:"index"`
	Project Project `json:"project" gorm:"foreignKey:UserID"`
}

type UserTags struct {
	TagID int `json:"tag_id" gorm:"index"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID    int       `json:"user_id" gorm:"index"`
}
