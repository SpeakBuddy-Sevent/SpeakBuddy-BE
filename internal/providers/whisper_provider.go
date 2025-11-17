package providers

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"time"
// )

// // WhisperProvider handles speech-to-text via OpenAI Whisper (Audio Transcriptions endpoint)
// type WhisperProvider struct {
// 	apiKey string
// }

// // NewWhisperProvider creates a new provider reading OPENAI_API_KEY from env
// func NewWhisperProvider() WhisperProvider {
// 	return WhisperProvider{
// 		apiKey: os.Getenv("OPENAI_API_KEY"),
// 	}
// }

// // Transcribe accepts an audioURL which can be:
// // - an http(s) URL (will be downloaded temporarily), or
// // - a local file path.
// // It returns the transcribed text (in plain string) or an error.
// func (w WhisperProvider) Transcribe(audioURL string) (string, error) {
// 	if w.apiKey == "" {
// 		return "", fmt.Errorf("OPENAI_API_KEY not set")
// 	}
// 	if audioURL == "" {
// 		return "", fmt.Errorf("audioURL is empty")
// 	}

// 	// determine whether audioURL is remote or local
// 	var localPath string
// 	var cleanup bool
// 	if strings.HasPrefix(strings.ToLower(audioURL), "http://") || strings.HasPrefix(strings.ToLower(audioURL), "https://") {
// 		// download file to temp
// 		tmp, err := downloadToTempFile(audioURL)
// 		if err != nil {
// 			return "", fmt.Errorf("failed to download audio: %w", err)
// 		}
// 		localPath = tmp
// 		cleanup = true
// 	} else {
// 		// assume local path
// 		localPath = audioURL
// 		cleanup = false
// 	}

// 	// ensure cleanup of temp file if created
// 	if cleanup {
// 		defer os.Remove(localPath)
// 	}

// 	// open file
// 	file, err := os.Open(localPath)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to open audio file: %w", err)
// 	}
// 	defer file.Close()

// 	// prepare multipart form
// 	var b bytes.Buffer
// 	writer := multipart.NewWriter(&b)

// 	// file part - field name is "file" per OpenAI audio transcription API
// 	part, err := writer.CreateFormFile("file", filepath.Base(localPath))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create form file: %w", err)
// 	}
// 	if _, err := io.Copy(part, file); err != nil {
// 		return "", fmt.Errorf("failed to copy file to form: %w", err)
// 	}

// 	// model param
// 	_ = writer.WriteField("model", "whisper-1")
// 	// language param - "id" for Bahasa Indonesia (you can set to "indonesian" if preferred)
// 	_ = writer.WriteField("language", "id")

// 	if err := writer.Close(); err != nil {
// 		return "", fmt.Errorf("failed to close multipart writer: %w", err)
// 	}

// 	// create request
// 	req, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", &b)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to create request: %w", err)
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	req.Header.Set("Authorization", "Bearer "+w.apiKey)

// 	// client with timeout
// 	client := &http.Client{Timeout: 120 * time.Second}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", fmt.Errorf("request error: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to read response: %w", err)
// 	}

// 	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
// 		return "", fmt.Errorf("whisper API error: status %d body: %s", resp.StatusCode, string(respBody))
// 	}

// 	// response is JSON like: { "text": "transcribed text..." }
// 	var parsed map[string]interface{}
// 	if err := json.Unmarshal(respBody, &parsed); err != nil {
// 		return "", fmt.Errorf("failed to parse whisper response: %w", err)
// 	}

// 	text, ok := parsed["text"].(string)
// 	if !ok {
// 		// fallback: maybe different field or structure
// 		return "", fmt.Errorf("unexpected whisper response format: %s", string(respBody))
// 	}

// 	return text, nil
// }

// // downloadToTempFile downloads a remote URL to a temp file and returns the local path.
// func downloadToTempFile(url string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", fmt.Errorf("failed http get: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
// 		return "", fmt.Errorf("download failed with status: %d", resp.StatusCode)
// 	}

// 	tmpFile, err := os.CreateTemp("", "audio-*"+filepath.Ext(url))
// 	if err != nil {
// 		return "", fmt.Errorf("failed create temp file: %w", err)
// 	}
// 	defer tmpFile.Close()

// 	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
// 		os.Remove(tmpFile.Name())
// 		return "", fmt.Errorf("failed write temp file: %w", err)
// 	}

// 	return tmpFile.Name(), nil
// }