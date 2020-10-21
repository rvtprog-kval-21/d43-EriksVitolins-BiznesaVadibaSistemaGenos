package migrations

import (
	"fmt"
	"golang-api/database"
	"golang-api/helpers/password"
	"golang-api/model"
)

func Migrate() {
	database.Open()
	userMigrate()
	database.Close()
}
func userMigrate() {
	err := database.DBConn.AutoMigrate(&model.User{})
	user := model.User{Email: "test@test.com", Password: password.HashAndSalt([]byte("test")), Role: "admin"}
	database.DBConn.Select("Email", "Password", "Role").Create(&user)
	if err != nil {
		panic("Migration failed")
	}
	fmt.Println(" ")
}
