package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang-api/controllers/user_controller"
	"golang-api/helpers/middleware"
)

func Init() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Post("/auth/login", user_controller.Login)

	api := app.Group("/api", middleware.Protected())

	admin := api.Group("/admin", func(context *fiber.Ctx) error {
		return middleware.IsAdmin(context)
	})
	fmt.Println(admin)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
