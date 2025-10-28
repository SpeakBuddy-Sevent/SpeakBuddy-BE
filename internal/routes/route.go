package routes

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/controllers"
)

type RouteSetup struct {
	// SessionController *controllers.SessionController
	SpeechController  *controllers.SpeechController
}

func NewRouteSetup(
	// sessionController *controllers.SessionController,
	speechController *controllers.SpeechController,
) *RouteSetup {
	return &RouteSetup{
		// SessionController: sessionController,
		SpeechController:  speechController,
	}
}


func (rs *RouteSetup) Setup(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	//api.Post("/sessions", controllers.CreateSession)
	//api.Get("/sessions/:id", controllers.GetSessionByID)

	api.Post("/speech/transcribe", rs.SpeechController.Transcribe)
}
