package migrations

import (
	"api/database"
	"api/model/blog"
	"api/model/calendar"
	"api/model/notifications"
	"api/model/projects"
	"api/model/timetable"
	"api/model/tracking"
	user2 "api/model/user"
	"api/utlis/password"
	"fmt"
	"time"
)

func Migrate() {
	database.Open()
	userMigrate()
}
func userMigrate() {
	migrateTables()
	populateTables()
}

func migrateTables() {
	err := database.DBConn.AutoMigrate(&user2.User{})
	err = database.DBConn.AutoMigrate(&projects.Project{})
	err = database.DBConn.AutoMigrate(&projects.Member{})
	err = database.DBConn.AutoMigrate(&projects.Tag{})
	err = database.DBConn.AutoMigrate(&projects.ProjectAnnouncements{})
	err = database.DBConn.AutoMigrate(&blog.Bloggers{})
	err = database.DBConn.AutoMigrate(&blog.Blogs{})
	err = database.DBConn.AutoMigrate(&notifications.Notifications{})
	err = database.DBConn.AutoMigrate(&calendar.Event{})
	err = database.DBConn.AutoMigrate(&calendar.EventMember{})
	err = database.DBConn.AutoMigrate(&timetable.Timetable{})
	err = database.DBConn.AutoMigrate(&user2.FollowingUser{})
	err = database.DBConn.AutoMigrate(&user2.AnnouncementsUser{})
	err = database.DBConn.AutoMigrate(&blog.BlogsLog{})
	err = database.DBConn.AutoMigrate(&tracking.Managers{})
	err = database.DBConn.AutoMigrate(&tracking.TrackedSubmission{})
	err = database.DBConn.AutoMigrate(&tracking.TrackedAttachment{})
	if err != nil {
		panic("Migration failed")
	}
	fmt.Println(" ")
}

func populateTables() {
	populateUser()
}

func populateUser() {
	user := user2.User{
		Email:       "test@test.com",
		Password:    password.HashAndSalt([]byte("test")),
		Role:        "admin",
		Avatar:      "/avatar/1/avatar.jpg",
		Name:        "Eriks",
		LastName:    "Vitolins",
		About:       "this is about me lol",
		Title:       "Ceo",
		Birthday:    time.Now(),
		NameDay:     time.Now(),
		PhoneNumber: "12355512",
		Background:  "/background/1/back.png",
	}
	database.DBConn.Select("Email", "Password", "Role", "Avatar", "Name", "LastName", "About", "Title", "Birthday", "NameDay", "PhoneNumber", "Background").Create(&user)
	user = user2.User{
		Email:       "test@test.lv",
		Password:    password.HashAndSalt([]byte("test")),
		Role:        "regular",
		Avatar:      "/avatar/2/avatar.jpg",
		Name:        "Laura",
		LastName:    "Vitolins",
		About:       "this is about me lol",
		Title:       "Qa",
		Birthday:    time.Now(),
		NameDay:     time.Now(),
		PhoneNumber: "12355512",
		Background:  "/background/2/back.png",
	}
	database.DBConn.Select("Email", "Password", "Role", "Avatar", "Name", "LastName", "About", "Title", "Birthday", "NameDay", "PhoneNumber", "Background").Create(&user)
	user = user2.User{
		Email:       "test@test.la",
		Password:    password.HashAndSalt([]byte("test")),
		Role:        "regular",
		Avatar:      "/avatar/3/avatar.jpg",
		Name:        "Siers",
		LastName:    "Vitolins",
		About:       "this is about me lol",
		Title:       "Automation Dev",
		Birthday:    time.Now(),
		NameDay:     time.Now(),
		PhoneNumber: "12355512",
		Background:  "/background/3/back.png",
	}
	database.DBConn.Select("Email", "Password", "Role", "Avatar", "Name", "LastName", "About", "Title", "Birthday", "NameDay", "PhoneNumber", "Background").Create(&user)
}
