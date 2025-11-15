package bootstrap

import (
	"os"

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
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/gen-lang-client-0235190640-7bd0c00a7ced.json")

	// Init database
	config.InitDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.SessionRecording{},
		&models.Feedback{},
	)

	app := fiber.New()

	// Providers
	whisperProvider := providers.NewWhisperProvider()
	googleSpeechProvider := providers.NewSpeechToTextProvider()
	geminiProvider := providers.NewGeminiProvider()

	// Repositories
	userRepo := repository.NewUserRepository(config.DB)
	feedbackRepo := repository.NewFeedbackRepository()
	sessionRepo := repository.NewSessionRepository(config.DB)

	// Services
	authService := services.NewAuthService(userRepo)
	feedbackService := services.NewFeedbackService(geminiProvider, feedbackRepo)
	sessionService := services.NewSessionService(googleSpeechProvider, geminiProvider, sessionRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	speechController := controllers.NewSpeechController(whisperProvider, googleSpeechProvider, sessionService)
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
