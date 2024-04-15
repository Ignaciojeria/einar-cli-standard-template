package subscription

import (
	"archetype/app/adapter/out/slog"
	"archetype/app/exception"
	"archetype/app/infrastructure/observability"
	"archetype/app/infrastructure/pubsubwrapper/subscriptionwrapper"
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

func (p messageProcessorStruct) subscriptionName() string {
	return "INSERT_YOUR_SUBSCRIPTION_NAME_HERE"
}

func (p messageProcessorStruct) Pull(ctx context.Context, m *pubsub.Message) (statusCode int, err error) {
	_, span := observability.Tracer.Start(ctx,
		"messageProcessorStruct",
		trace.WithSpanKind(trace.SpanKindConsumer), trace.WithAttributes(
			attribute.String("subscription.name", p.subscriptionName()),
			attribute.String("message.id", m.ID),
			attribute.String("message.publishTime", m.PublishTime.String()),
		))

	var dataModel interface{}
	defer func() {
		statusCode = subscriptionwrapper.HandleMessageAcknowledgement(span,
			&subscriptionwrapper.HandleMessageAcknowledgementDetails{
				SubscriptionName: p.subscriptionName(),
				Error:            err,
				Message:          m,
				ErrorsRequiringNack: []error{
					exception.INTERNAL_SERVER_ERROR,
					exception.EXTERNAL_SERVER_ERROR,
					exception.HTTP_NETWORK_ERROR,
					exception.PUBSUB_BROKER_ERROR,
				},
				CustomLogFields: slog.CustomLogFields{
					"customIndexField": "MyCustomFieldForIndexWhenSearchLogs",
				},
			})
		span.End()
	}()
	if err := json.Unmarshal(m.Data, &dataModel); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

/* ----- Default Initialization & Configuration Settings ----- */

type messageProcessorStruct struct {
	subscriptionReference *pubsub.Subscription
}

func init() {
	ioc.Registry(
		newMessageProcessor,
		subscriptionwrapper.NewSubscriptionManager)
}

func newMessageProcessor(
	subscriptionManager subscriptionwrapper.SubscriptionManager) subscriptionwrapper.MessageProcessor {
	messageProcessor := messageProcessorStruct{}
	subscriptionRef := subscriptionManager.Subscription(messageProcessor.subscriptionName())
	subscriptionRef.ReceiveSettings.MaxOutstandingMessages = 5
	messageProcessor.subscriptionReference = subscriptionRef
	sm := subscriptionManager.
		WithMessageProcessor(messageProcessor).
		WithPushHandler("/subscription/" + messageProcessor.subscriptionName())
	go sm.Start()
	return messageProcessor
}

func (p messageProcessorStruct) SubscriptionRef() *pubsub.Subscription {
	return p.subscriptionReference
}
