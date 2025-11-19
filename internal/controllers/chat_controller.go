package controllers

import (
	"time"
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
    userID, ok := ctx.Locals("user_id").(uint)
    if !ok {
        return ctx.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }

    therapistID := ctx.Params("therapistID")
    if therapistID == "" {
        return ctx.Status(400).JSON(fiber.Map{"error": "therapistID required"})
    }

    var req request.SendMessageRequest
    if err := ctx.BodyParser(&req); err != nil {
        return ctx.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    if req.Text == "" {
        return ctx.Status(400).JSON(fiber.Map{"error": "text is required"})
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
        ID:        msg.ID.Hex(),
        ChatID:    msg.ChatID.Hex(),
        SenderID:  msg.SenderID,
        Text:      msg.Text,
        Timestamp: msg.Timestamp.UTC().Format(time.RFC3339),
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
			ID:        m.ID.Hex(),
			ChatID:    m.ChatID.Hex(),
			SenderID:  m.SenderID,
			Text:      m.Text,
			Timestamp: m.Timestamp.UTC().Format(time.RFC3339),
		})
	}

	return ctx.JSON(res)
}

func (cc *ChatController) MyChats(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)
	userIDStr := fmt.Sprint(userID)

	list, err := cc.service.GetMyChats(fmt.Sprint(userIDStr))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	var res []response.ChatResponse
	for _, c := range list {
		therapistID := ""
        for _, p := range c.Participants {
            if p != userIDStr {
                therapistID = p
            }
        }
		// Participants left as-is (string array)
		lastTime := ""
		if !c.LastMessageTime.IsZero() {
			lastTime = c.LastMessageTime.UTC().Format(time.RFC3339)
		}
		res = append(res, response.ChatResponse{
			ID:              c.ID.Hex(),
			Participants:    c.Participants,
			TherapistID:     therapistID,
			LastMessageText: c.LastMessageText,
			LastMessageTime: lastTime,
		})
	}

	return ctx.JSON(fiber.Map{
        "data": res,          
        "message": "success", 
    })
}

func (cc *ChatController) SendToChat(c *fiber.Ctx) error {
    chatID := c.Params("chatID")
    if chatID == "" {
        return c.Status(400).JSON(fiber.Map{"error": "chatID required"})
    }

    rawUserID := c.Locals("user_id")
    if rawUserID == nil {
        return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
    }

    userID := fmt.Sprint(rawUserID)

    var req request.SendMessageRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
    }

    if req.Text == "" {
        return c.Status(400).JSON(fiber.Map{"error": "text is required"})
    }

    msg, err := cc.service.SendMessageToChat(chatID, userID, req)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(response.MessageResponse{
        ID:        msg.ID.Hex(),
        ChatID:    msg.ChatID.Hex(),
        SenderID:  msg.SenderID,
        Text:      msg.Text,
        Timestamp: msg.Timestamp.UTC().Format(time.RFC3339),
    })
}



