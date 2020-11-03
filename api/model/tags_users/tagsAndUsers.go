package tags_users

import "api/model/user"

type TagsAndUsers struct {
	User   user.User `gorm:"foreignKey:UserId"`
	UserId int       `json:"owner"`
	Tag    user.User `gorm:"foreignKey:TagId"`
	TagId  int       `json:"owner"`
}
