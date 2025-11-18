package controllers

import (
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/request"
	"speakbuddy/pkg/dto/response"

	"fmt"
)

type ChatController struct {
	service services.ChatService
}

func NewChatController(service services.ChatService) *ChatController {
	return &ChatController{service}
}

func (cc *ChatController) SendMessage(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	therapistID := ctx.Params("therapistID")

	var req request.SendMessageRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	msg, err := cc.service.SendMessage(
		fmt.Sprint(userID),
		therapistID,
		req,
	)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(response.MessageResponse{
		ID:        msg.ID,
		ChatID:    msg.ChatID,
		SenderID:  msg.SenderID,
		Text:      msg.Text,
		Timestamp: msg.Timestamp.String(),
	})
}

func (cc *ChatController) GetMessages(ctx *fiber.Ctx) error {
	chatID := ctx.Params("chatID")

	list, err := cc.service.GetChatMessages(chatID)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var res []response.MessageResponse
	for _, m := range list {
		res = append(res, response.MessageResponse{
			ID:        m.ID,
			ChatID:    m.ChatID,
			SenderID:  m.SenderID,
			Text:      m.Text,
			Timestamp: m.Timestamp.String(),
		})
	}

	return ctx.JSON(res)
}

func (cc *ChatController) MyChats(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	list, err := cc.service.GetMyChats(fmt.Sprint(userID))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var res []response.ChatResponse
	for _, c := range list {
		res = append(res, response.ChatResponse{
			ID:              c.ID,
			Participants:    c.Participants,
			LastMessageText: c.LastMessageText,
			LastMessageTime: c.LastMessageTime.String(),
		})
	}

	return ctx.JSON(res)
}
