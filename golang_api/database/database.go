package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Open() {
	var err error
	dbAddress := "genos:mua3Rgadasdadasdasdasasagas2a41aASeXK8XMUPvMN@tcp(127.0.0.1:3306)/genos?charset=utf8mb4&parseTime=True&loc=Local"
	DBConn, err = gorm.Open(mysql.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
		fmt.Println(DBConn)
	}
}
func Close() {
	db, err := DBConn.DB()
	if err != nil {
		panic("Failed to connect to database")
	} else {
		db.Close()
	}
}
