package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/config"
	"speakbuddy/internal/models"
	"speakbuddy/internal/routes"
	"speakbuddy/internal/providers"
	"speakbuddy/internal/controllers"
)

func InitializeApp() *fiber.App {
	config.InitDB()

	config.DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Feedback{})

	app := fiber.New()

	whisper := providers.NewWhisperProvider()

	speechController := controllers.NewSpeechController(whisper)

	rs := routes.NewRouteSetup(speechController)
	rs.Setup(app)

	return app
}
