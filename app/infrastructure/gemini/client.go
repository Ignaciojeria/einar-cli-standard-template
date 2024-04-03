package gemini

import (
	"archetype/app/configuration"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func init() {
	ioc.Registry(NewClient, configuration.NewConf)
}

func NewClient(conf configuration.Conf) (*genai.Client, error) {
	return genai.NewClient(context.Background(), option.WithAPIKey(conf.GeminiApiKey))
}
