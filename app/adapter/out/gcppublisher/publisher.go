package gcppublisher

import (
	"archetype/app/shared/constants"
	"archetype/app/shared/infrastructure/gcppubsub"
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/logging"
	"archetype/app/shared/systemerr"
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type INewPublishEvent func(ctx context.Context, input interface{}) error

func init() {
	ioc.Registry(
		NewPublishEvent,
		gcppubsub.NewClient,
		logging.NewLogger)
}
func NewPublishEvent(c *pubsub.Client, logger logging.Logger) INewPublishEvent {
	topicName := "INSERT_YOUR_TOPIC_NAME_HERE"
	topic := c.Topic(topicName)
	return func(ctx context.Context, input interface{}) error {
		_, span := observability.Tracer.Start(ctx, "INewPublishEvent",
			trace.WithSpanKind(trace.SpanKindProducer),
			trace.WithAttributes(attribute.String(constants.TopicName, topicName)),
		)
		defer span.End()

		bytes, err := json.Marshal(input)
		if err != nil {
			return err
		}

		message := &pubsub.Message{
			Attributes: map[string]string{
				"customAttribute1": "attr1",
				"customAttribute2": "attr2",
			},
			Data: bytes,
		}

		result := topic.Publish(ctx, message)
		// Get the server-generated message ID.
		messageID, err := result.Get(ctx)

		if err != nil {
			span.SetStatus(codes.Error, systemerr.PUBSUB_BROKER_ERROR.Error())
			span.RecordError(err)
			logger.SpanLogger(span).Error(
				systemerr.PUBSUB_BROKER_ERROR.Error(),
				constants.TopicName, topicName,
				constants.Error, err.Error())
			return systemerr.PUBSUB_BROKER_ERROR
		}

		span.SetStatus(codes.Ok, "Message published with ID: "+messageID)

		return nil
	}
}
