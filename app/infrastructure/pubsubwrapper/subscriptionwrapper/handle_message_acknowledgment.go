package subscriptionwrapper

import (
	"archetype/app/adapter/out/slog"
	"archetype/app/constants"
	"errors"

	"cloud.google.com/go/pubsub"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type HandleMessageAcknowledgementDetails struct {
	SubscriptionName    string
	Error               error
	Message             *pubsub.Message
	ErrorsRequiringNack []error
	CustomLogFields     map[string]interface{}
}

func HandleMessageAcknowledgement(span trace.Span, details *HandleMessageAcknowledgementDetails) {
	if details.Error != nil {
		span.RecordError(details.Error)
		span.SetStatus(codes.Error, details.Error.Error())
		slog.SpanLogger(span).Error(
			details.SubscriptionName+"_exception",
			subscription_name, details.SubscriptionName,
			constants.Fields, details.CustomLogFields,
			constants.Error, details.Error,
		)
		for _, err := range details.ErrorsRequiringNack {
			if errors.Is(details.Error, err) {
				span.AddEvent("Event processing nacked, retrying",
					trace.WithAttributes(attribute.String(constants.Error, details.Error.Error())))
				details.Message.Nack()
				return
			}
		}
		span.AddEvent("Event discarded",
			trace.WithAttributes(attribute.String(constants.Error, details.Error.Error())))
		details.Message.Ack()
		return
	}
	span.SetStatus(codes.Ok, details.SubscriptionName+"_succedded")
	slog.SpanLogger(span).Info(
		details.SubscriptionName+"_succedded",
		subscription_name, details.SubscriptionName,
		constants.Fields, details.CustomLogFields,
	)
	details.Message.Ack()
}
