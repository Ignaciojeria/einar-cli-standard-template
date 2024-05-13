package subscriptionwrapper

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type MessageProcessor func(ctx context.Context, m *pubsub.Message) (statusCode int, err error)
