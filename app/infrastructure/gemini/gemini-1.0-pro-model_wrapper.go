package gemini

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/google/generative-ai-go/genai"
)

type gemini1Dot0ProModelWrapper struct {
	*genai.GenerativeModel
}

func init() {
	ioc.Registry(newGemini1Dot0ProModelWrapper, newClient)
}
func newGemini1Dot0ProModelWrapper(client *genai.Client) gemini1Dot0ProModelWrapper {
	return gemini1Dot0ProModelWrapper{
		GenerativeModel: client.GenerativeModel("gemini-1.0-pro"),
	}
}

func Gemini1Dot0ProModelWrapper() gemini1Dot0ProModelWrapper {
	return ioc.Get[gemini1Dot0ProModelWrapper](newGemini1Dot0ProModelWrapper)
}
