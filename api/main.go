package main

import (
	"api/database"
	"api/database/migrations"
	"api/router"
	"api/services/online"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "server" {
		go online.ClearOldOnes()
		database.Open()
		router.Init()
	} else if argsWithoutProg[0] == "migrate" {
		migrations.Migrate()
	}
}
