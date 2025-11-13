package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"strconv"
)

type FeedbackController struct {
	service services.FeedbackService
}

func NewFeedbackController(service services.FeedbackService) *FeedbackController {
	return &FeedbackController{service: service}
}

func (fc *FeedbackController) AnalyzeFeedback(c *fiber.Ctx) error {
	sessionIDStr := c.FormValue("session_id")
	text := c.FormValue("text")

	if text == "" || sessionIDStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "session_id dan text wajib diisi",
		})
	}

	sessionID, err := strconv.ParseUint(sessionIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "session_id tidak valid",
		})
	}

	result, err := fc.service.AnalyzeAndSaveFeedback(uint(sessionID), text)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}
