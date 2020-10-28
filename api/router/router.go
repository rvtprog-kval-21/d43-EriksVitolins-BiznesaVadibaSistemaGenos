package router

import (
	"api/controllers/user"
	"api/helpers/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"net/http"
)

func Init() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(":8000")
}

func setupRoutes(app *fiber.App) {
	routeMiddlewareStart(app)

	authRoutes(app)

	api := app.Group("/api", middleware.Protected())
	admin := api.Group("/admin", func(context *fiber.Ctx) error {
		return middleware.IsAdmin(context)
	})

	adminRoutes(admin)
	apiRoutes(api)

	routeMiddlewareAfter(app)
}

func authRoutes(app *fiber.App) {
	app.Post("/auth/login", user.Login)
}

func adminRoutes(admin fiber.Router) {
	admin.Post("/users", user.Index)
	admin.Post("/user/:id/lock", user.LockUser)
	admin.Post("/user/:id/unlock", user.UnlockUser)
}

func apiRoutes(api fiber.Router) {
	api.Get("/user/:id/profile", user.User)
}
func routeMiddlewareStart(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Authorization, Content-Length, X-Requested-With, Accept",
		AllowMethods:     "GET, POST, OPTIONS",
		AllowCredentials: true,
	}))
	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(logger.New())

}

func routeMiddlewareAfter(app *fiber.App) {
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("api/storage"),
	}))
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

}
