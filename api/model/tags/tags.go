package tags

import (
	"api/database"
	"api/model/user"
	"gorm.io/gorm"
	"strconv"
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

func GetAllMemberTags(id interface{}, isPublic bool) []Tag {
	var members []Member
	var tags []Tag
	database.DBConn.Where("user_id = ?", id).Find(&members)

	data := NameList(members)
	if isPublic {
		database.DBConn.Preload("Members.User").Where("id in ?", data).Where("is_public = ?", true).Find(&tags)
	} else {
		database.DBConn.Preload("Members.User").Where("id in ?", data).Find(&tags)
	}
	return tags
}

func NameList(u []Member) []int {
	var list []int
	for _, iter := range u {
		list = append(list, iter.TagID)
	}
	return list
}

func Profile(id string, userID interface{}) (Tag, bool) {
	var tag Tag
	var members Member
	isMember := false
	database.DBConn.Where("user_id = ?", userID).Where("tag_id = ?", id).Find(&members)
	database.DBConn.Preload("Members.User").Find(&tag, "id = ?", id)
	if members.TagID == 0 {
		if tag.IsPublic == false {
			return Tag{}, isMember
		}
	} else {
		isMember = true
	}

	return tag, isMember
}

func GetTagAndUser(id string, userID interface{}) (Tag, Member) {
	var tag Tag
	var members Member
	database.DBConn.Where("user_id = ?", userID).Where("tag_id = ?", id).Find(&members)
	database.DBConn.Preload("Members.User").Find(&tag, "id = ?", id)

	return tag, members
}

func JoinTag(id string, userID interface{}) interface{} {
	tag, isMember := Profile(id, userID)
	if isMember {
		return "Already a member"
	} else if tag.ID == 0 {
		return "No such tag available"
	}
	intID, _ := strconv.Atoi(id)

	user := Member{
		IsAdmin: false,
		IsOwner: false,
		UserID:  int(userID.(float64)),
		TagID:   intID,
	}
	results := database.DBConn.Create(&user)
	return results.Error
}

func DeleteTag(id string, userID interface{}) interface{} {
	tag, member := GetTagAndUser(id, userID)
	if tag.ID == 0 {
		return "Tag doesn't exist"
	} else if member.IsOwner == false {
		return "Only the owner of the tag can delete a tag"
	}
	results := database.DBConn.Unscoped().Delete(&tag)
	if results.Error != nil {
		return results.Error
	}
	results = database.DBConn.Where("tag_id = ?", id).Delete(Member{})
	return results.Error
}

func UpdateName(id string, userID interface{}, data string) interface{} {
	tag, member := GetTagAndUser(id, userID)
	var temp Tag
	if tag.ID == 0 {
		return "Tag doesn't exist"
	} else if member.IsAdmin == false {
		return "Only and admin of the tag can change a tag"
	}
	tag.Name = data
	database.DBConn.Where("name", data).First(&temp)
	if temp.ID != 0 {
		return "Tag name already exists"
	}
	results := database.DBConn.Model(&Tag{}).Where("id = ?", id).Update("name", data)
	return results.Error
}

func UpdateAbout(id string, userID interface{}, data string) interface{} {
	tag, member := GetTagAndUser(id, userID)
	if tag.ID == 0 {
		return "Tag doesn't exist"
	} else if member.IsAdmin == false {
		return "Only and admin of the tag can change a tag"
	}
	tag.About = data
	results := database.DBConn.Model(&Tag{}).Where("id = ?", id).Update("about", data)
	return results.Error
}

func UpdatePublic(id string, userID interface{}, isPublic bool) interface{} {
	tag, member := GetTagAndUser(id, userID)
	if tag.ID == 0 {
		return "Tag doesn't exist"
	} else if member.IsAdmin == false {
		return "Only and admin of the tag can change a tag"
	}
	results := database.DBConn.Model(&Tag{}).Where("id = ?", id).Update("is_public", isPublic)
	return results.Error
}
