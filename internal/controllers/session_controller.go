package controllers

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"speakbuddy/pkg/dto/request"
// 	"speakbuddy/internal/services"
// )

// func CreateSession(c *fiber.Ctx) error {
// 	var req request.CreateSessionRequest

// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	session, err := services.CreateSession(req)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(session)
// }

// func GetSessionByID(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	session, err := services.GetSessionByID(id)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"error": "Not found"})
// 	}
// 	return c.JSON(session)
// }