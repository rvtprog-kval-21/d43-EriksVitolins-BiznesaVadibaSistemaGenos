package server

import (
	"golang-api/controllers/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":8000")
}

func setupRoutes(app *fiber.App) {
	app.Post("/auth/login", user.Login)

	api := app.Group("/api", middleware.Protected())

	admin := api.Group("/admin", func(context *fiber.Ctx) error {
		return middleware.IsAdmin(context)
	})

	admin.Post("/users", user.Index)
	api.Get("/user/:id/profile", user.User)
	admin.Post("/user/:id/lock", user.LockUser)
	admin.Post("/user/:id/unlock", user.UnlockUser)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}

func middleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Authorization, Content-Length, X-Requested-With, Accept",
		AllowMethods:     "GET, POST, OPTIONS",
		AllowCredentials: true,
	}))
	app.Use(logger.New())
}
