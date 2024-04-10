package pubsubwrapper

import (
	"context"

	"cloud.google.com/go/pubsub"
)

type MessageProcessor interface {
	Pull(ctx context.Context, m *pubsub.Message) (statusCode int, err error)
	SubscriptionRef() *pubsub.Subscription
}
