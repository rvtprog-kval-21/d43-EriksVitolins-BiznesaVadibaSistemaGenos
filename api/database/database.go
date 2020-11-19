package database

import (
	"api/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConn database connection
var DBConn *gorm.DB

// Open Opens database connection
func Open() {
	fmt.Println("opened")
	var err error
	dbAddress := config.DbUser + ":" + config.DbPassword + "@tcp(" + config.DbHost + ")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err = gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
		fmt.Println(DBConn)
	}
}

// Close Closes database connection
func Close() {
	fmt.Println("closed")
	db, err := DBConn.DB()
	if err != nil {
		panic("Failed to connect to database")
	} else {
		db.Close()
	}
}
