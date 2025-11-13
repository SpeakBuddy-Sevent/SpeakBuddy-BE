package controllers

import (
	"os"
	"github.com/gofiber/fiber/v2"
	"speakbuddy/internal/providers"
)

type SpeechController struct {
	Whisper providers.WhisperProvider
}

func NewSpeechController(provider providers.WhisperProvider) *SpeechController {
	return &SpeechController{Whisper: provider,}
}

func (c *SpeechController) Transcribe(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no audio file uploaded",
		})
	}

	tmpPath := "./tmp/" + file.Filename
	if err := ctx.SaveFile(file, tmpPath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save file",
		})
	}
	defer os.Remove(tmpPath) // ngeremove path temporary di akhir (defer)

	// speech to text buat mp3?
	text, err := c.Whisper.Transcribe(tmpPath)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// harapan return string -> bener salah aman
	return ctx.JSON(fiber.Map{
		"text": text,
	})
}
