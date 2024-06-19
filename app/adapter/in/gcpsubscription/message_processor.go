package gcpsubscription

import (
	"archetype/app/shared/infrastructure/gcppubsub/subscriptionwrapper"
	"archetype/app/shared/infrastructure/observability"
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	ioc.Registry(
		newMessageProcessor,
		subscriptionwrapper.NewSubscriptionManager)
}
func newMessageProcessor(
	sm subscriptionwrapper.SubscriptionManager,
) subscriptionwrapper.MessageProcessor {
	subscriptionName := "INSERT_YOUR_SUBSCRIPTION_NAME_HERE"
	subscriptionRef := sm.Subscription(subscriptionName)
	subscriptionRef.ReceiveSettings.MaxOutstandingMessages = 5
	messageProcessor := func(ctx context.Context, m *pubsub.Message) (int, error) {
		_, span := observability.Tracer.Start(ctx,
			"messageProcessorStruct",
			trace.WithSpanKind(trace.SpanKindConsumer), trace.WithAttributes(
				attribute.String("subscription.name", subscriptionRef.String()),
				attribute.String("message.id", m.ID),
				attribute.String("message.publishTime", m.PublishTime.String()),
			))
		defer span.End()
		var input interface{}
		if err := json.Unmarshal(m.Data, &input); err != nil {
			span.SetStatus(codes.Error, err.Error())
			m.Ack()
			return http.StatusAccepted, err
		}
		m.Ack()
		return http.StatusOK, nil
	}
	go sm.WithMessageProcessor(messageProcessor).
		WithPushHandler("/subscription/" + subscriptionName).
		Start(subscriptionRef)
	return messageProcessor
}
