package gemini

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/google/generative-ai-go/genai"
)

func init() {
	ioc.Registry(NewGemini1Dot0ProModel, NewClient)
}

func NewGemini1Dot0ProModel(client *genai.Client) *genai.GenerativeModel {
	return client.GenerativeModel("gemini-1.0-pro")
}
