package migrations

import (
	"api/database"
	"api/model/tags"
	user2 "api/model/user"
	"api/utlis/password"
	"fmt"
)

func Migrate() {
	database.Open()
	userMigrate()
	database.Close()
}
func userMigrate() {
	err := database.DBConn.AutoMigrate(&user2.User{})
	err = database.DBConn.AutoMigrate(&tags.Tags{})
	user := user2.User{Email: "test@test.com", Password: password.HashAndSalt([]byte("test")), Role: "admin"}
	database.DBConn.Select("Email", "Password", "Role").Create(&user)
	if err != nil {
		panic("Migration failed")
	}
	fmt.Println(" ")
}
