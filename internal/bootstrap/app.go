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
	// Init database
	config.InitDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Feedback{},
	)

	app := fiber.New()

	// Providers
	whisperProvider := providers.NewWhisperProvider()
	geminiProvider := providers.NewGeminiProvider()

	// Repositories
	userRepo := repository.NewUserRepository(config.DB)
	feedbackRepo := repository.NewFeedbackRepository()

	// Services
	authService := services.NewAuthService(userRepo)
	feedbackService := services.NewFeedbackService(geminiProvider, feedbackRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	speechController := controllers.NewSpeechController(whisperProvider)
	feedbackController := controllers.NewFeedbackController(feedbackService)

	// Routes
	rs := routes.NewRouteSetup(
		authController,
		speechController,
		feedbackController,
	)

	rs.Setup(app)

	return app
}
