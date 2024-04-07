package chat

import (
	"archetype/app/infrastructure/gemini"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type IChat interface {
	SendMessage(ctx context.Context, domain interface{}) (interface{}, error)
}

type chat struct {
	model gemini.Gemini1Dot0ProModelWrapper
}

func init() {
	ioc.Registry(NewChat, gemini.NewGemini1Dot0ProModelWrapper)
}
func NewChat(model gemini.Gemini1Dot0ProModelWrapper) IChat {
	return chat{
		model: model,
	}
}

func (s chat) SendMessage(ctx context.Context, domain interface{}) (interface{}, error) {
	// design your custom prompt and chat here using gemini-1.0-pro model
	return s.model.EphemeralChatExpectJSONResult(ctx, `
	Write a JSON object in a single line.

	Here's an example format for your request:
	{"Message":"Hello World"}`)
}

func Instance() IChat {
	return ioc.Get[IChat](NewChat)
}
