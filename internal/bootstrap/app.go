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
	config.InitMongo()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.DataAnak{},
		&models.Feedback{},
		&models.ReadingExerciseTemplate{},
		&models.ExerciseItem{},
		&models.ExerciseAttempt{},
		&models.Consultation{},
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
	profileRepo := repository.NewProfileRepository(config.DB)
	dataAnakRepo := repository.NewDataAnakRepository(config.DB)

	feedbackRepo := repository.NewFeedbackRepository()
	templateRepo := repository.NewReadingExerciseTemplateRepository(config.DB)
	itemRepo := repository.NewExerciseItemRepository(config.DB)
	attemptRepo := repository.NewExerciseAttemptRepository(config.DB)
	consultationRepo := repository.NewConsultationRepository(config.DB)
	chatRepo := repository.NewChatRepository(config.MongoDB)

	// Services
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	profileService := services.NewProfileService(profileRepo)
	dataAnakService := services.NewDataAnakService(dataAnakRepo)

	feedbackService := services.NewFeedbackService(geminiProvider, feedbackRepo)
	exerciseService := services.NewExerciseService(googleSpeechProvider, geminiProvider, attemptRepo, itemRepo, templateRepo)
	consultationService := services.NewConsultationService(consultationRepo)
	chatService := services.NewChatService(chatRepo)
	therapistService := services.NewTherapistService(userRepo)

	// Controllers
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController(userService)
	profileController := controllers.NewProfileController(profileService)
	dataAnakController := controllers.NewDataAnakController(dataAnakService)

	feedbackController := controllers.NewFeedbackController(feedbackService)
	exerciseController := controllers.NewExerciseController(exerciseService)
	consultationController := controllers.NewConsultationController(consultationService)
	chatController := controllers.NewChatController(chatService)
	therapistController := controllers.NewTherapistController(therapistService)

	// Routes
	rs := routes.NewRouteSetup(
		authController,
		feedbackController,
		exerciseController,
		profileController,
		dataAnakController,
		userController,
		consultationController,
		chatController,
		therapistController,
	)

	rs.Setup(app)

	return app
}
