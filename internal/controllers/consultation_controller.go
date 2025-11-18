package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/pkg/dto/response"
	"strconv"
)

type ConsultationController struct {
	service services.ConsultationService
}

func NewConsultationController(service services.ConsultationService) *ConsultationController {
	return &ConsultationController{service}
}

func (cc *ConsultationController) Book(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	therapistIDParam := ctx.Params("therapistUserID")
	therapistID64, _ := strconv.ParseUint(therapistIDParam, 10, 32)
	therapistUserID := uint(therapistID64)

	var req request.CreateConsultationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body",
		})
	}

	result, err := cc.service.BookConsultation(userID, therapistUserID, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(response.ConsultationResponse{
	ID:           result.ID,
	UserID:       result.UserID,
	TherapistUserID:  result.TherapistUserID,
	Date: result.Date.Format("2006-01-02 15:04"),
	IsPaid:       result.IsPaid,
})

}

func (cc *ConsultationController) MyConsultations(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	list, err := cc.service.GetMyConsultations(userID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var res []response.ConsultationResponse
	for _, c := range list {
		res = append(res, response.ConsultationResponse{
		ID:           c.ID,
		UserID:       c.UserID,
		TherapistUserID:  c.TherapistUserID,
		Date: c.Date.Format("2006-01-02 15:04"),
		IsPaid:       c.IsPaid,
	})
	}

	return ctx.JSON(res)
}

func (cc *ConsultationController) TherapistConsultations(ctx *fiber.Ctx) error {
	therapistID := ctx.Locals("user_id").(uint) // therapist login

	list, err := cc.service.GetTherapistConsultations(therapistID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var res []response.ConsultationResponse
	for _, c := range list {
		res = append(res, response.ConsultationResponse{
		ID:           c.ID,
		UserID:       c.UserID,
		TherapistUserID:  c.TherapistUserID,
		Date: c.Date.Format("2006-01-02 15:04"),
		IsPaid:       c.IsPaid,
	})
	}

	return ctx.JSON(res)
}
