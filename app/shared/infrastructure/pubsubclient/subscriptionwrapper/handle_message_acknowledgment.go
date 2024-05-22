package subscriptionwrapper

import (
	"archetype/app/shared/constants"
	"archetype/app/shared/logger"
	"errors"
	"net/http"

	"cloud.google.com/go/pubsub"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type HandleMessageAcknowledgement func(span trace.Span, details *HandleMessageAcknowledgementDetails) int

type HandleMessageAcknowledgementDetails struct {
	SubscriptionName    string
	Error               error
	Message             *pubsub.Message
	ErrorsRequiringNack []error
	CustomLogFields     logger.CustomLogFields
}

func init() {
	ioc.Registry(
		NewHandleMessageAcknowledgement,
		logger.NewLogger)
}
func NewHandleMessageAcknowledgement(l logger.Logger) HandleMessageAcknowledgement {
	return func(span trace.Span, details *HandleMessageAcknowledgementDetails) int {
		if details.Error != nil {
			span.RecordError(details.Error)
			span.SetStatus(codes.Error, details.Error.Error())
			l.SpanLogger(span).Error(
				details.SubscriptionName+"_exception",
				subscription_name, details.SubscriptionName,
				constants.Fields, details.CustomLogFields,
				constants.Error, details.Error,
				constants.MessageAttributes, details.Message.Attributes,
				constants.MessageData, string(details.Message.Data),
			)
			for _, err := range details.ErrorsRequiringNack {
				if errors.Is(details.Error, err) {
					span.AddEvent("Event processing nacked, retrying",
						trace.WithAttributes(attribute.String(constants.Error, details.Error.Error())))
					details.Message.Nack()
					return http.StatusInternalServerError
				}
			}
			span.AddEvent("Event discarded",
				trace.WithAttributes(attribute.String(constants.Error, details.Error.Error())))
			details.Message.Ack()
			return http.StatusAccepted
		}
		span.SetStatus(codes.Ok, details.SubscriptionName+"_succedded")
		l.SpanLogger(span).Info(
			details.SubscriptionName+"_succedded",
			subscription_name, details.SubscriptionName,
			constants.Fields, details.CustomLogFields,
			constants.MessageAttributes, details.Message.Attributes,
			constants.MessageData, string(details.Message.Data),
		)
		details.Message.Ack()
		return http.StatusOK
	}
}