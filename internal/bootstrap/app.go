package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/config"
	"speakbuddy/internal/models"
	"speakbuddy/internal/routes"
)

func InitializeApp() *fiber.App {
	config.InitDB()

	config.DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Feedback{})

	app := fiber.New()

	routes.SetupRoutes(app)

	return app
}
