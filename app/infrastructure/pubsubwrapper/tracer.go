package pubsubwrapper

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("pubsub-subscription")
