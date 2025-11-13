package bootstrap

import (
	"speakbuddy/config"
	"speakbuddy/internal/controllers"
	"speakbuddy/internal/models"
	"speakbuddy/internal/providers"
	"speakbuddy/internal/repository"
	"speakbuddy/internal/routes"
	"speakbuddy/internal/services"

	"github.com/gofiber/fiber/v2"
)

func InitializeApp() *fiber.App {
	config.InitDB()

	config.DB.AutoMigrate(&models.User{}, &models.Session{}, &models.Feedback{})

	app := fiber.New()

	whisperProvider := providers.NewWhisperProvider()
	geminiProvider := providers.NewGeminiProvider()

	speechController := controllers.NewSpeechController(whisperProvider)

	feedbackRepo := repository.NewFeedbackRepository()
	feedbackService := services.NewFeedbackService(geminiProvider, feedbackRepo)
	feedbackController := controllers.NewFeedbackController(feedbackService)

	rs := routes.NewRouteSetup(speechController, feedbackController)
	rs.Setup(app)

	return app
}
