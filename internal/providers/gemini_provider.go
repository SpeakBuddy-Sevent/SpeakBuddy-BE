package providers

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type GeminiProvider interface {
	GetFeedbackFromGemini(inputText string) (string, error)
}

type geminiProvider struct {
	client *genai.Client
}

func NewGeminiProvider() GeminiProvider {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY is not set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize Gemini client: %v", err))
	}

	return &geminiProvider{client: client}
}

func (g *geminiProvider) GetFeedbackFromGemini(inputText string) (string, error) {
	ctx := context.Background()
	prompt := fmt.Sprintf(
		"Analisis pelafalan anak berdasarkan teks berikut dan berikan saran perbaikan yang singkat dan ramah: %s",
		inputText,
	)

	result, err := g.client.Models.GenerateContent(ctx, "gemini-2.5-flash", genai.Text(prompt), nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate feedback: %v", err)
	}

	text := result.Text()
	if text == "" {
		text = "Tidak ada respons dari model Gemini"
	}

	return text, nil
}
