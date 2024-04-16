package publisher

import (
	"archetype/app/adapter/out/slog"
	"archetype/app/constants"
	"archetype/app/exception"
	"archetype/app/infrastructure/observability"
	"archetype/app/infrastructure/pubsubwrapper/topicwrapper"
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func PublishEvent(ctx context.Context, REPLACE_BY_YOUR_DOMAIN interface{}) (err error) {
	topicName := "INSERT YOUR TOPIC NAME HERE"

	_, span := observability.Tracer.Start(ctx, "PublishEvent",
		trace.WithSpanKind(trace.SpanKindProducer),
		trace.WithAttributes(attribute.String(constants.TopicName, topicName)),
	)
	defer span.End()

	bytes, err := json.Marshal(REPLACE_BY_YOUR_DOMAIN)
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

	result := topicwrapper.Get(topicName).Publish(ctx, message)
	// Get the server-generated message ID.
	messageID, err := result.Get(ctx)

	if err != nil {
		span.SetStatus(codes.Error, exception.PUBSUB_BROKER_ERROR.Error())
		span.RecordError(err)
		slog.
			LogSpanError(span, exception.PUBSUB_BROKER_ERROR.Error(),
				slog.CustomLogFields{
					constants.Error: err.Error(),
				})
		return exception.PUBSUB_BROKER_ERROR
	}

	span.SetStatus(codes.Ok, "Message published with ID: "+messageID)

	return nil
}
