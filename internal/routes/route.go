package routes

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/controllers"
)

type RouteSetup struct {
	// SessionController *controllers.SessionController
	SpeechController  *controllers.SpeechController
	FeedbackController *controllers.FeedbackController
}

func NewRouteSetup(
	// sessionController *controllers.SessionController,
	speechController *controllers.SpeechController,
	feedbackController *controllers.FeedbackController,
) *RouteSetup {
	return &RouteSetup{
		// SessionController: sessionController,
		SpeechController:  speechController,
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

	//api.Post("/sessions", controllers.CreateSession)
	//api.Get("/sessions/:id", controllers.GetSessionByID)

	api.Post("/speech/transcribe", rs.SpeechController.Transcribe)

	api.Post("/feedback/analyze", rs.FeedbackController.AnalyzeFeedback)

}
