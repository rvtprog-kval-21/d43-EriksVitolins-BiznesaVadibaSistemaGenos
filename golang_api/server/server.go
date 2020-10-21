package server

import (
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
	app.Get("/apa", user_controller.Signup)
	api.Get("/user/:id/logout", user_controller.Logout)
}
