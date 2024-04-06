package usecase

import (
	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/google/generative-ai-go/genai"
)

func init() {
	ioc.Registry(NewChatSession)
}

type ChatSession struct {
	*genai.GenerativeModel
}

func NewChatSession() ChatSession {
	return ChatSession{}
}
