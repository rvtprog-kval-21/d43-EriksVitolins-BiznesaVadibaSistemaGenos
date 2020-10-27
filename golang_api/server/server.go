package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"golang-api/controllers/user_controller"
	"golang-api/helpers/middleware"
	"net/http"
)

func Init() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":8000")
}

func setupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Authorization, Content-Length, X-Requested-With, Accept",
		AllowMethods:     "GET, POST, OPTIONS",
		AllowCredentials: true,
	}))
	app.Use(logger.New())

	app.Post("/auth/login", user_controller.Login)

	api := app.Group("/api", middleware.Protected())
	admin := api.Group("/admin", func(context *fiber.Ctx) error {
		return middleware.IsAdmin(context)
	})
	admin.Post("/users", user_controller.Index)
	api.Get("/user/:id/profile", user_controller.User)
	admin.Post("/user/:id/lock", user_controller.LockUser)
	admin.Post("/user/:id/unlock", user_controller.UnlockUser)

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("golang_api/storage"),
	}))
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
