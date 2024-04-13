package firestore_repository

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("firestore_repository")
