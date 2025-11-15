package providers

import (
	"context"
	//"io"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
	"google.golang.org/api/option"
)

type SpeechToTextProvider struct {
	client *speech.Client
}

func NewSpeechToTextProvider() *SpeechToTextProvider {
	credentialsPath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")

	ctx := context.Background()
	var client *speech.Client
	var err error

	if credentialsPath != "" {
		client, err = speech.NewClient(ctx, option.WithCredentialsFile(credentialsPath))
	} else {
		client, err = speech.NewClient(ctx)
	}

	if err != nil {
		panic("Failed to initialize Google Speech-to-Text client: " + err.Error())
	}

	return &SpeechToTextProvider{client: client}
}

func (p *SpeechToTextProvider) TranscribeAudio(audio []byte) (string, error) {
	ctx := context.Background()

	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:        speechpb.RecognitionConfig_MP3,
			LanguageCode:    "id-ID",
			SampleRateHertz: 44100,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: audio,
			},
		},
	}

	resp, err := p.client.Recognize(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Results) == 0 {
		return "", nil
	}

	return resp.Results[0].Alternatives[0].Transcript, nil
}
