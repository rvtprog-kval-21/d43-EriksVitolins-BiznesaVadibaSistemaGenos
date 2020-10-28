package main

import (
	"api/database/migrations"
	"api/router"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "server" {
		router.Init()
	} else if argsWithoutProg[0] == "migrate" {
		migrations.Migrate()
	}
}
