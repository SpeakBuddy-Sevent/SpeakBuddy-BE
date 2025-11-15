package controllers

// import (
// 	"os"
// 	"speakbuddy/internal/providers"
// 	"speakbuddy/internal/services"
// 	"speakbuddy/pkg/dto/response"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// )

// type SpeechController struct {
// 	Whisper        providers.WhisperProvider
// 	GoogleSpeech   *providers.SpeechToTextProvider
// 	SessionService services.SessionService
// }

// func NewSpeechController(
// 	whisperProvider providers.WhisperProvider,
// 	googleSpeechProvider *providers.SpeechToTextProvider,
// 	sessionService services.SessionService,
// ) *SpeechController {
// 	return &SpeechController{
// 		Whisper:        whisperProvider,
// 		GoogleSpeech:   googleSpeechProvider,
// 		SessionService: sessionService,
// 	}
// }

// func (c *SpeechController) Transcribe(ctx *fiber.Ctx) error {
// 	file, err := ctx.FormFile("file")
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "no audio file uploaded",
// 		})
// 	}

// 	tmpPath := "./tmp/" + file.Filename
// 	if err := ctx.SaveFile(file, tmpPath); err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to save file",
// 		})
// 	}
// 	defer os.Remove(tmpPath)

// 	audioBytes, err := os.ReadFile(tmpPath)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to read audio file",
// 		})
// 	}

// 	text, err := c.GoogleSpeech.TranscribeAudio(audioBytes)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "transcription failed: " + err.Error(),
// 		})
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"text": text,
// 	})
// }

// func (c *SpeechController) CreateSession(ctx *fiber.Ctx) error {
// 	userIDInterface := ctx.Locals("user_id")
// 	if userIDInterface == nil {
// 		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": "user_id not found in token",
// 		})
// 	}

// 	userID := uint(userIDInterface.(float64))

// 	targetText := ctx.FormValue("target_text")
// 	if targetText == "" {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "target_text wajib diisi",
// 		})
// 	}

// 	session, err := c.SessionService.CreateSession(userID, targetText)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to create session: " + err.Error(),
// 		})
// 	}

// 	return ctx.JSON(fiber.Map{
// 		"id":          session.ID,
// 		"user_id":     session.UserID,
// 		"target_text": session.TargetText,
// 		"created_at":  session.CreatedAt,
// 	})
// }

// func (c *SpeechController) TranscribeAndAnalyze(ctx *fiber.Ctx) error {
// 	file, err := ctx.FormFile("file")
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "no audio file uploaded",
// 		})
// 	}

// 	sessionIDStr := ctx.FormValue("session_id")
// 	if sessionIDStr == "" {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "session_id wajib diisi",
// 		})
// 	}

// 	sessionID, err := strconv.ParseUint(sessionIDStr, 10, 32)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "invalid session_id format",
// 		})
// 	}

// 	tmpPath := "./tmp/" + file.Filename
// 	if err := ctx.SaveFile(file, tmpPath); err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to save file",
// 		})
// 	}
// 	defer os.Remove(tmpPath)

// 	audioBytes, err := os.ReadFile(tmpPath)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to read audio file",
// 		})
// 	}

// 	// Transcribe dan analyze, return SessionRecording
// 	recording, err := c.SessionService.TranscribeAndAnalyze(uint(sessionID), audioBytes)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "analysis failed: " + err.Error(),
// 		})
// 	}

// 	// Get session untuk ambil target text
// 	session, err := c.SessionService.GetSessionWithRecordings(uint(sessionID))
// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "failed to get session",
// 		})
// 	}

// 	res := response.TranscribeAndAnalyzeResponse{
// 		TranscribedText: recording.TranscribedText,
// 		TargetText:      session.TargetText,
// 		RecordingID:     recording.ID,
// 		Feedback:        recording.AIFeedback,
// 		AIModel:         recording.AIModel,
// 		Accuracy:        recording.Accuracy,
// 	}

// 	return ctx.JSON(res)
// }
