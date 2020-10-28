package migrations

import (
	"api/database"
	user "api/model"
	"api/utlis/password"
	"fmt"
)

func Migrate() {
	database.Open()
	userMigrate()
	database.Close()
}
func userMigrate() {
	err := database.DBConn.AutoMigrate(&user.User{})
	user := user.User{Email: "test@test.com", Password: password.HashAndSalt([]byte("test")), Role: "admin"}
	database.DBConn.Select("Email", "Password", "Role").Create(&user)
	if err != nil {
		panic("Migration failed")
	}
	fmt.Println(" ")
}
