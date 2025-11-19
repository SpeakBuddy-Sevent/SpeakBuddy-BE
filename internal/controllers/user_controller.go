package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/request"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (uc *UserController) UpdateName(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	var req request.UpdateUserNameRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := uc.service.UpdateName(userID, req.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "name updated",
	})
}

func (uc *UserController) FindByID(ctx *fiber.Ctx) error {
    userID := ctx.Locals("user_id").(uint)

    user, err := uc.service.GetByID(userID)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return ctx.JSON(user)
}

