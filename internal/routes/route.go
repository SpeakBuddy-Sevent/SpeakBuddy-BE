package routes

import (
	"speakbuddy/internal/controllers"
	"speakbuddy/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type RouteSetup struct {
	AuthController     *controllers.AuthController
	FeedbackController *controllers.FeedbackController
	ExerciseController *controllers.ExerciseController
}

func NewRouteSetup(
	authController *controllers.AuthController,
	feedbackController *controllers.FeedbackController,
	exerciseController *controllers.ExerciseController,
) *RouteSetup {
	return &RouteSetup{
		AuthController:     authController,
		FeedbackController: feedbackController,
		ExerciseController: exerciseController,
	}
}

func (rs *RouteSetup) Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:3000",
        AllowMethods:     "GET,POST,PUT,PATCH,DELETE,OPTIONS",
        AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
        ExposeHeaders:    "Content-Length",
        AllowCredentials: true,
    }))

	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api.Post("/auth/register", rs.AuthController.Register)
	api.Post("/auth/login", rs.AuthController.Login)

	protected := api.Group("/", middleware.AuthRequired)
	{
		protected.Post("/feedback/analyze", rs.FeedbackController.AnalyzeFeedback)

		// Exercise endpoints
		protected.Get("/exercise/levels", rs.ExerciseController.GetLevels)
		protected.Post("/exercise/start", rs.ExerciseController.StartExercise)
		protected.Get("/exercise/:exerciseID/next-item", rs.ExerciseController.GetNextItem)
		protected.Post("/exercise/record", rs.ExerciseController.RecordAttempt)
	}

}
