package pubsub

import (
	"archetype/app/configuration"
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

type ClientWrapper struct {
	*pubsub.Client
}

func init() {
	ioc.Registry(NewClientWrapper, configuration.NewConf)
}
func NewClientWrapper(conf configuration.Conf) (ClientWrapper, error) {
	c, err := pubsub.NewClient(context.Background(), conf.GOOGLE_PROJECT_ID)
	if conf.GOOGLE_PROJECT_ID == "" {
		return ClientWrapper{}, errors.New("GOOGLE_PROJECT_ID is not present")
	}
	if err != nil {
		return ClientWrapper{}, err
	}
	return ClientWrapper{
		Client: c,
	}, nil
}
