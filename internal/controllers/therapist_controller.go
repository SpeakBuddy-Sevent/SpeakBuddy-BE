package controllers

import (
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/response"

	"github.com/gofiber/fiber/v2"
)

type TherapistController struct {
	service services.TherapistService
}

func NewTherapistController(s services.TherapistService) *TherapistController {
	return &TherapistController{service: s}
}

func (tc *TherapistController) GetAll(ctx *fiber.Ctx) error {
	therapists, err := tc.service.GetAllTherapists()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var res []response.TherapistResponse
	for _, t := range therapists {
		res = append(res, response.TherapistResponse{
			ID:    t.ID,
			Name:  t.Name,
			Email: t.Email,
		})
	}

	return ctx.JSON(res)
}

func (tc *TherapistController) GetByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid therapist id",
		})
	}

	therapist, err := tc.service.GetTherapistByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "therapist not found",
		})
	}

	return ctx.JSON(response.TherapistResponse{
		ID:    therapist.ID,
		Name:  therapist.Name,
		Email: therapist.Email,
	})
}
