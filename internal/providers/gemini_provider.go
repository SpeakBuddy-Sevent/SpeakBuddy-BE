package providers

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type GeminiProvider interface {
	GetFeedbackFromGemini(targetText, inputText string) (string, error)
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

func (g *geminiProvider) GetFeedbackFromGemini(targetText, inputText string) (string, error) {
	ctx := context.Background()
	prompt := fmt.Sprintf(`
Seorang anak sedang belajar berbicara dan melafalkan kata dengan bantuan aplikasi terapi bicara.

Kalimat yang seharusnya diucapkan: "%s"
Kalimat yang diucapkan anak: "%s"

Tugasmu adalah membantu anak dengan cara:
1. Berikan pujian atau semangat di awal agar anak merasa percaya diri.
2. Jelaskan dengan lembut bagian mana yang terdengar kurang tepat (jangan gunakan kata "salah" secara langsung).
3. Berikan contoh pelafalan yang benar, bisa dibagi per suku kata bila perlu.
4. Tutup dengan kalimat penyemangat singkat yang positif.

Gunakan gaya berbicara yang hangat, sederhana, dan cocok untuk anak-anak dengan speech delay.
Pastikan hasilmu mudah dimengerti ketika diubah menjadi suara oleh sistem text-to-speech.
`, targetText, inputText)

	result, err := g.client.Models.GenerateContent(ctx, "gemini-2.0-flash", genai.Text(prompt), nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate feedback: %v", err)
	}

	text := result.Text()
	if text == "" {
		text = "Tidak ada respons dari model Gemini"
	}

	return text, nil
}
