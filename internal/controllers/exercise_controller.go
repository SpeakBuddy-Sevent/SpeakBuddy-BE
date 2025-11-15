package controllers

import (
	"os"
	"speakbuddy/internal/services"
	"speakbuddy/pkg/dto/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ExerciseController struct {
	exerciseService services.ExerciseService
}

func NewExerciseController(exerciseService services.ExerciseService) *ExerciseController {
	return &ExerciseController{
		exerciseService: exerciseService,
	}
}

// RecordAttempt - user record audio untuk 1 soal, transcribe + analyze langsung
func (ec *ExerciseController) RecordAttempt(ctx *fiber.Ctx) error {
	// Parse form data
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "no audio file uploaded",
		})
	}

	// Get item_id dari form
	itemIDStr := ctx.FormValue("item_id")
	if itemIDStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "item_id wajib diisi",
		})
	}

	itemID, err := strconv.ParseUint(itemIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid item_id format",
		})
	}

	// Save file temporarily
	tmpPath := "./tmp/" + file.Filename
	if err := ctx.SaveFile(file, tmpPath); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save file",
		})
	}
	defer os.Remove(tmpPath)

	// Read audio file
	audioBytes, err := os.ReadFile(tmpPath)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to read audio file",
		})
	}

	// Transcribe + Analyze
	attempt, err := ec.exerciseService.TranscribeAndAnalyzeAttempt(uint(itemID), audioBytes)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "transcription/analysis failed: " + err.Error(),
		})
	}

	res := response.ExerciseAttemptResponse{
		ID:              attempt.ID,
		TranscribedText: attempt.TranscribedText,
		Feedback:        attempt.AIFeedback,
		AIModel:         attempt.AIModel,
		Accuracy:        attempt.Accuracy,
		CreatedAt:       attempt.CreatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
