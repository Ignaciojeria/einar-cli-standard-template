package pubsubwrapper

import (
	"archetype/app/shared/configuration"
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

type ClientWrapper struct {
	client *pubsub.Client
}

func init() {
	ioc.Registry(NewClientWrapper, configuration.NewConf)
}

func NewClientWrapper(conf configuration.Conf) (ClientWrapper, error) {
	c, err := pubsub.NewClient(context.Background(), conf.GOOGLE_PROJECT_ID)
	if conf.GOOGLE_PROJECT_ID == "" {
		return ClientWrapper{
			client: &pubsub.Client{},
		}, errors.New("GOOGLE_PROJECT_ID is not present")
	}
	if err != nil {
		return ClientWrapper{
			client: &pubsub.Client{},
		}, err
	}
	return ClientWrapper{
		client: c,
	}, nil
}

func (cw *ClientWrapper) Client() *pubsub.Client {
	return cw.client
}

func (cw ClientWrapper) Subscription(id string) *pubsub.Subscription {
	return cw.client.Subscription(id)
}