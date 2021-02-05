package main

import (
	"api/database"
	"api/database/migrations"
	"api/router"
	"api/services/online"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if argsWithoutProg[0] == "server" {
		go online.Initiate()
		database.Open()
		app := router.Init()
		server := &http.Server{
			Addr:           ":8000",
			Handler:        app,
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   5 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(server.ListenAndServe())
	} else if argsWithoutProg[0] == "migrate" {
		migrations.Migrate()
	}
}
