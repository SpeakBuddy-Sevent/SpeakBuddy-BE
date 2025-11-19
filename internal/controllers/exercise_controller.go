package controllers

import (
	"os"
	"speakbuddy/internal/models"
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
	// Get user_id dari token
	userIDInterface := ctx.Locals("user_id")
	if userIDInterface == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "user_id not found in token",
		})
	}
	userID := uint(userIDInterface.(float64))

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
	// tmpPath := "./tmp/" + file.Filename
	// if err := ctx.SaveFile(file, tmpPath); err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "failed to save file",
	// 	})
	// }
	// defer os.Remove(tmpPath)

	tmpFile, err := os.CreateTemp("/tmp", "upload-*.wav")
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "failed to create temp file",
        })
    }
    defer os.Remove(tmpFile.Name())

    // Simpan file
    if err := ctx.SaveFile(file, tmpFile.Name()); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "failed to save file",
        })
    }

	audioBytes, err := os.ReadFile(tmpFile.Name())
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "failed to read audio file",
        })
    }


	// Read audio file
	// audioBytes, err := os.ReadFile(tmpPath)
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": "failed to read audio file",
	// 	})
	// }

	// Transcribe + Analyze
	attempt, err := ec.exerciseService.TranscribeAndAnalyzeAttempt(userID, uint(itemID), audioBytes)
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

// GetLevels - get semua exercise levels
func (ec *ExerciseController) GetLevels(ctx *fiber.Ctx) error {
	exercises, err := ec.exerciseService.GetAllExerciseTemplates()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get levels: " + err.Error(),
		})
	}

	// Extract hanya level + title
	var levels []fiber.Map
	for _, ex := range exercises {
		levels = append(levels, fiber.Map{
			"id":    ex.ID,
			"title": ex.Title,
			"level": ex.Level,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(levels)
}

// StartExercise - user pilih level dan mulai exercise, return first item
func (ec *ExerciseController) StartExercise(ctx *fiber.Ctx) error {
	level := ctx.FormValue("level")
	if level == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "level wajib diisi (dasar|menengah|lanjut)",
		})
	}

	exercise, err := ec.exerciseService.GetExerciseByLevel(level)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get exercise: " + err.Error(),
		})
	}

	if exercise.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "exercise level tidak ditemukan",
		})
	}

	// Get first item
	var firstItem *models.ExerciseItem
	if len(exercise.Items) > 0 {
		firstItem = &exercise.Items[0]
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"exercise_id": exercise.ID,
		"level":       exercise.Level,
		"title":       exercise.Title,
		"first_item": fiber.Map{
			"id":          firstItem.ID,
			"item_number": firstItem.ItemNumber,
			"target_text": firstItem.TargetText,
		},
	})
}

// GetNextItem - get soal berikutnya dari exercise
func (ec *ExerciseController) GetNextItem(ctx *fiber.Ctx) error {
	exerciseIDStr := ctx.Params("exerciseID")
	currentItemNumberStr := ctx.Query("current_item_number")

	if exerciseIDStr == "" || currentItemNumberStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "exercise_id dan current_item_number wajib diisi",
		})
	}

	exerciseID, err := strconv.ParseUint(exerciseIDStr, 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid exercise_id format",
		})
	}

	currentItemNumber, err := strconv.Atoi(currentItemNumberStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid current_item_number format",
		})
	}

	nextItem, err := ec.exerciseService.GetNextItem(uint(exerciseID), currentItemNumber)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get next item: " + err.Error(),
		})
	}

	if nextItem == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "semua soal sudah selesai",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":          nextItem.ID,
		"item_number": nextItem.ItemNumber,
		"target_text": nextItem.TargetText,
	})
}
