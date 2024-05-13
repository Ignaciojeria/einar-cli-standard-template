package subscription

import (
	"archetype/app/shared/exception"
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/infrastructure/pubsubclient/subscriptionwrapper"
	"archetype/app/shared/logger"
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	ioc.Registry(
		newMessageProcessor,
		subscriptionwrapper.NewSubscriptionManager)
}
func newMessageProcessor(sm subscriptionwrapper.SubscriptionManager) subscriptionwrapper.MessageProcessor {
	subscriptionName := "INSERT_YOUR_SUBSCRIPTION_NAME_HERE"
	subscriptionRef := sm.Subscription(subscriptionName)
	subscriptionRef.ReceiveSettings.MaxOutstandingMessages = 5
	messageProcessor := func(ctx context.Context, m *pubsub.Message) (statusCode int, err error) {
		_, span := observability.Tracer.Start(ctx,
			"messageProcessorStruct",
			trace.WithSpanKind(trace.SpanKindConsumer), trace.WithAttributes(
				attribute.String("subscription.name", subscriptionRef.String()),
				attribute.String("message.id", m.ID),
				attribute.String("message.publishTime", m.PublishTime.String()),
			))
		var input interface{}
		defer func() {
			statusCode = subscriptionwrapper.HandleMessageAcknowledgement(span,
				&subscriptionwrapper.HandleMessageAcknowledgementDetails{
					SubscriptionName: subscriptionRef.String(),
					Error:            err,
					Message:          m,
					ErrorsRequiringNack: []error{
						exception.INTERNAL_SERVER_ERROR,
						exception.EXTERNAL_SERVER_ERROR,
						exception.HTTP_NETWORK_ERROR,
						exception.PUBSUB_BROKER_ERROR,
					},
					CustomLogFields: logger.CustomLogFields{
						"customIndexField": "MyCustomFieldForIndexWhenSearchLogs",
					},
				})
			span.End()
		}()
		if err := json.Unmarshal(m.Data, &input); err != nil {
			return statusCode, err
		}
		return statusCode, nil
	}
	go sm.WithMessageProcessor(messageProcessor).
		WithPushHandler("/subscription/" + subscriptionName).
		Start(subscriptionRef)
	return messageProcessor
}
