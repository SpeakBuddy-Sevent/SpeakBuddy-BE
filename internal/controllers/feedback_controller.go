package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/pkg/dto/response"
)

type FeedbackController struct {
	service services.FeedbackService
}

func NewFeedbackController(service services.FeedbackService) *FeedbackController {
	return &FeedbackController{service: service}
}

func (fc *FeedbackController) AnalyzeFeedback(c *fiber.Ctx) error {
	var req request.FeedbackRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "request body tidak valid",
		})
	}

	if req.SessionID == 0 || req.TargetText == "" || req.InputText == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "session_id dan text wajib diisi",
		})
	}

	result, err := fc.service.AnalyzeAndSaveFeedback(req.SessionID, req.TargetText, req.InputText)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	res := response.FeedbackResponse{
		ID:        result.ID,
		SessionID: result.SessionID,
		AIModel:   result.AIModel,
		Feedback:  result.Feedback,
		CreatedAt: result.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
