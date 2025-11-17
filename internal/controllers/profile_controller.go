package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/pkg/dto/response"
)

type ProfileController struct {
	service services.ProfileService
}

func NewProfileController(service services.ProfileService) *ProfileController {
	return &ProfileController{service}
}

func (pc *ProfileController) Upsert(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	var req request.CreateProfileRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	prof, err := pc.service.CreateOrUpdateProfile(userID, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(response.ProfileResponse{
		ID:     prof.ID,
		UserID: prof.UserID,
		Age:    prof.Age,
		Sex:    prof.Sex,
		Phone:  prof.Phone,
	})
}

func (pc *ProfileController) Get(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	prof, err := pc.service.GetProfile(userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}

	return ctx.JSON(response.ProfileResponse{
		ID:     prof.ID,
		UserID: prof.UserID,
		Age:    prof.Age,
		Sex:    prof.Sex,
		Phone:  prof.Phone,
	})
}
