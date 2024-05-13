package pubsubclient

import (
	"archetype/app/shared/configuration"
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
)

func init() {
	ioc.Registry(pubsub.NewClient, configuration.NewConf)
}

func NewClient(conf configuration.Conf) (*pubsub.Client, error) {
	c, err := pubsub.NewClient(context.Background(), conf.GOOGLE_PROJECT_ID)
	if conf.GOOGLE_PROJECT_ID == "" {
		return &pubsub.Client{}, errors.New("GOOGLE_PROJECT_ID is not present")
	}
	if err != nil {
		return &pubsub.Client{}, err
	}
	return c, nil
}
