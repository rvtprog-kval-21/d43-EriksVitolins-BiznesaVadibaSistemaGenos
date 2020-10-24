package main

import (
	"golang-api/database/migrations"
	"golang-api/server"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "server" {
		server.Init()
	} else if argsWithoutProg[0] == "migrate" {
		migrations.Migrate()
	}
}
