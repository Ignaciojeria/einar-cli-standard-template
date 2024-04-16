package firestore_repository

import (
	"archetype/app/shared/infrastructure/firebasewrapper/firestorewrapper"
	"archetype/app/shared/infrastructure/observability"
	"context"

	"cloud.google.com/go/firestore"
	"go.opentelemetry.io/otel/trace"
)

func RunFirestoreOperation(ctx context.Context, entity interface{}) error {
	_, span := observability.Tracer.Start(ctx,
		"RunFirestoreOperation",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	var _ *firestore.CollectionRef = firestorewrapper.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
	//Do something with collection.
	return nil
}
