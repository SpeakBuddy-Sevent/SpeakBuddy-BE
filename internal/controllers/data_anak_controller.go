package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/pkg/dto/response"
)

type DataAnakController struct {
	service services.DataAnakService
}

func NewDataAnakController(service services.DataAnakService) *DataAnakController {
	return &DataAnakController{service}
}

func (dc *DataAnakController) Upsert(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	var req request.CreateDataAnakRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	data, err := dc.service.CreateOrUpdate(userID, req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(response.DataAnakResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		ChildName: data.ChildName,
		ChildAge:  data.ChildAge,
		ChildSex:  data.ChildSex,
	})
}

func (dc *DataAnakController) Get(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	data, err := dc.service.Get(userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}

	return ctx.JSON(response.DataAnakResponse{
		ID:        data.ID,
		UserID:    data.UserID,
		ChildName: data.ChildName,
		ChildAge:  data.ChildAge,
		ChildSex:  data.ChildSex,
	})
}
