package chat

import (
	"archetype/app/infrastructure/gemini"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type IChat interface {
	EphemeralChatExpectJSONResult(ctx context.Context, domain interface{}) (interface{}, error)
}

type chat struct {
	model gemini.Gemini1Dot0ProModelWrapper
}

func init() {
	ioc.Registry(newChat, gemini.NewGemini1Dot0ProModelWrapper)
}
func newChat(model gemini.Gemini1Dot0ProModelWrapper) IChat {
	return chat{
		model: model,
	}
}

func (s chat) EphemeralChatExpectJSONResult(ctx context.Context, domain interface{}) (interface{}, error) {
	return s.model.EphemeralChatExpectJSONResult(ctx, `
	Write a JSON object in a single line.

	Here's an example format for your request:
	{"Message":"Hello World"}`)
}

func Chat() IChat {
	return ioc.Get[IChat](newChat)
}
