package bootstrap

import (
	"os"

	"speakbuddy/config"
	"speakbuddy/internal/controllers"
	"speakbuddy/internal/models"
	"speakbuddy/internal/providers"
	"speakbuddy/internal/repository"
	"speakbuddy/internal/routes"
	"speakbuddy/internal/seeder"
	"speakbuddy/internal/services"

	"github.com/gofiber/fiber/v2"
)

func InitializeApp() *fiber.App {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/gen-lang-client-0235190640-7bd0c00a7ced.json")

	// Init database
	config.InitDB()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Feedback{},
		&models.ReadingExerciseTemplate{},
		&models.ExerciseItem{},
		&models.ExerciseAttempt{},
	)

	// Seed default exercises
	if err := seeder.SeedExercises(config.DB); err != nil {
		panic("seeder error: " + err.Error())
	}

	app := fiber.New()

	// Providers
	googleSpeechProvider := providers.NewSpeechToTextProvider()
	geminiProvider := providers.NewGeminiProvider()

	// Repositories
	userRepo := repository.NewUserRepository(config.DB)
	feedbackRepo := repository.NewFeedbackRepository()
	itemRepo := repository.NewExerciseItemRepository(config.DB)
	attemptRepo := repository.NewExerciseAttemptRepository(config.DB)

	// Services
	authService := services.NewAuthService(userRepo)
	feedbackService := services.NewFeedbackService(geminiProvider, feedbackRepo)
	exerciseService := services.NewExerciseService(googleSpeechProvider, geminiProvider, attemptRepo, itemRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	feedbackController := controllers.NewFeedbackController(feedbackService)
	exerciseController := controllers.NewExerciseController(exerciseService)

	// Routes
	rs := routes.NewRouteSetup(
		authController,
		feedbackController,
		exerciseController,
	)

	rs.Setup(app)

	return app
}
