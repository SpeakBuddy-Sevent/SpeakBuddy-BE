package routes

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api.Post("/sessions", controller.CreateSession)
	api.Get("/sessions/:id", controller.GetSessionByID)
}
