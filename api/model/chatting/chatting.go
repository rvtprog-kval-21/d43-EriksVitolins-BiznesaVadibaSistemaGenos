package chatting

import (
	"api/database"
	"api/model/user"
	"time"
)

type Rooms struct {
	ID              int                `gorm:"primaryKey;not null" json:"id"`
	About           string             `json:"about"`
	Name            string             `json:"name"`
	Avatar          string             `json:"avatar"`
	Participants    []RoomParticipants `json:"participants"`
	UpdatedAt       *time.Time         `json:"updated_at"`
	Messages        []RoomMessages     `json:"messages"`
	NotSeenMessages []MessageViews     `json:"not_seen_messages"`
}

type RoomParticipants struct {
	ID        int       `gorm:"primaryKey;not null" json:"id"`
	Rooms     Rooms     `json:"rooms" gorm:"foreignKey:RoomsID"`
	RoomsID   int       `json:"rooms_id" gorm:"index"`
	User      user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID    int       `json:"user_id" gorm:"index"`
	IsAdmin   bool      `json:"is_admin"`
}

type RoomMessages struct {
	ID             int       `gorm:"primaryKey;not null" json:"id"`
	Sent           time.Time `json:"sent"`
	Rooms          Rooms     `json:"rooms" gorm:"foreignKey:RoomsID"`
	RoomsID        int       `json:"rooms_id" gorm:"index"`
	User           user.User `json:"user" gorm:"foreignKey:UserID"`
	UserID         int       `json:"user_id" gorm:"index"`
	Message        string    `json:"message"`
	IsDeleted      bool      `json:"is_deleted"`
	IsNotification bool      `json:"is_notification" `
	ReferenceID    int       `json:"Reference_id;omitempty" gorm:"index"`
	Attachment     string    `json:"attachment;omitempty"`
	Audio          string    `json:"audio;omitempty"`
}

type MessageViews struct {
	ID        int          `gorm:"primaryKey;not null" json:"id"`
	Seen      bool         `json:"seen" gorm:"index"`
	Message   RoomMessages `json:"message" gorm:"foreignKey:MessageID"`
	MessageID int          `json:"message_id" gorm:"index"`
	User      user.User    `json:"user" gorm:"foreignKey:UserID"`
	UserID    int          `json:"user_id" gorm:"index"`
	Rooms     Rooms        `json:"rooms" gorm:"foreignKey:RoomsID"`
	RoomsID   int          `json:"rooms_id" gorm:"index"`
}

func SaveRoom(obj *Rooms) {
	database.DBConn.Create(&obj)
}

func SaveView(obj MessageViews) {
	database.DBConn.Create(&obj)
}

func SaveParticipants(obj []RoomParticipants) {
	database.DBConn.Create(&obj)
}

func DeleteRoom(obj Rooms) {
	database.DBConn.Create(&obj)
}

func GetRooms(id interface{}) []Rooms {
	var rooms []Rooms
	subquey := database.DBConn.Select("rooms_id").Where("user_id = ?", id).Table("room_participants")
	database.DBConn.Preload("NotSeenMessages", "seen = 0 AND user_id = ?", id).Preload("Participants.User").Preload("Messages").Where("id in (?)", subquey).Where("is_deleted = ?", 0).Find(&rooms)
	return rooms
}

func GetRoom(id interface{}, idRoom int) Rooms {
	var rooms Rooms
	subquey := database.DBConn.Select("rooms_id").Where("user_id = ?", id).Where("rooms_id = ?", idRoom).Table("room_participants")
	database.DBConn.Preload("Participants.User").Preload("Messages.User").Where("id in (?)", subquey).Where("is_deleted = ?", 0).Find(&rooms)
	return rooms
}

func SaveMessage(obj RoomMessages) (RoomMessages, []RoomParticipants) {
	var arr []RoomParticipants
	database.DBConn.Create(&obj)
	database.DBConn.Where("rooms_id = ?", obj.RoomsID).Find(&arr)

	return obj, arr
}

func UpdateAt(roomId int)  {
	database.DBConn.Model(&Rooms{}).Where("id = ?", roomId).Update("updated_at", time.Now())
}

func GetUnViewedMessages(userID int, roomID int) []RoomMessages {
	var arr []RoomMessages
	subquey := database.DBConn.Select("message_id").Where("user_id = ?", userID).Where("rooms_id = ?", roomID).Where("seen = ?", 0).Table("message_views")
	database.DBConn.Preload("User").Where("id in (?)", subquey).Find(&arr)
	return arr
}

func UpdateViews(userID int, ids []int) {
	database.DBConn.Model(&MessageViews{}).Where("user_id = ?", userID).Where("message_id IN (?)", ids).Update("seen", 1)
}

func GetMember(roomID int, userID interface{}) RoomParticipants{
	var usr RoomParticipants
	database.DBConn.Where("rooms_id = ?", roomID).Where("user_id = ?", userID).Find(&usr)
	return usr
}

func UpdateName(rooms Rooms) {
	database.DBConn.Model(&Rooms{}).Where("id = ?", rooms.ID).Update("name", rooms.Name)
}

func UpdateAbout(rooms Rooms) {
	database.DBConn.Model(&Rooms{}).Where("id = ?", rooms.ID).Update("about", rooms.About)
}

func UpdateAvatar(rooms Rooms) {
	database.DBConn.Model(&Rooms{}).Where("id = ?", rooms.ID).Update("avatar", rooms.Avatar)
}

func GetNonMembers(roomId int) []user.User{
	var users []user.User
	database.DBConn.Raw("SELECT DISTINCT users.email, users.id FROM `users` left join room_participants on room_participants.user_id = users.id WHERE users.id NOT IN (SELECT room_participants.user_id FROM room_participants WHERE room_participants.rooms_id = ?)", roomId).Scan(&users)
	return users
}

func LeaveAt(roomId int, id interface{}) {
	database.DBConn.Where("user_id = ?", id).Where("rooms_id = ?", roomId).Delete(&RoomParticipants{})
}

func DeleteMessage(obj RoomMessages) {
	database.DBConn.Model(&obj).Where("id = ?", obj.ID).Update("is_deleted", 1)
}

func SaveNotification(obj RoomMessages) {
		database.DBConn.Select("rooms_id", "is_notification", "message", "sent").Create(&obj)
}

func GetUnreadCount(id interface{}) int64 {
	var count int64
	subquey := database.DBConn.Select("rooms_id").Where("user_id = ?", id).Table("room_participants")
	database.DBConn.Model(&MessageViews{}).Where("rooms_id in (?)", subquey).Where("user_id = ?", id).Where("seen = ?", 0).Count(&count)
	return count
}