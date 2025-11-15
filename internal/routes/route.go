package routes

import (
	"speakbuddy/internal/controllers"
	"speakbuddy/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteSetup struct {
	AuthController *controllers.AuthController
	// SessionController *controllers.SessionController
	SpeechController   *controllers.SpeechController
	FeedbackController *controllers.FeedbackController
}

func NewRouteSetup(
	authController *controllers.AuthController,
	// sessionController *controllers.SessionController,
	speechController *controllers.SpeechController,
	feedbackController *controllers.FeedbackController,
) *RouteSetup {
	return &RouteSetup{
		AuthController: authController,
		// SessionController: sessionController,
		SpeechController:   speechController,
		FeedbackController: feedbackController,
	}
}

func (rs *RouteSetup) Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"testing api": "mantap"})
	})

	api.Post("/auth/register", rs.AuthController.Register)
	api.Post("/auth/login", rs.AuthController.Login)
	//api.Post("/speech/transcribe-and-analyze", rs.SpeechController.TranscribeAndAnalyze) // for testing without auth

	protected := api.Group("/", middleware.AuthRequired)
	{
		protected.Post("/speech/create-session", rs.SpeechController.CreateSession)
		protected.Post("/speech/transcribe", rs.SpeechController.Transcribe)
		protected.Post("/speech/transcribe-and-analyze", rs.SpeechController.TranscribeAndAnalyze)

		protected.Post("/feedback/analyze", rs.FeedbackController.AnalyzeFeedback)
	}

}
