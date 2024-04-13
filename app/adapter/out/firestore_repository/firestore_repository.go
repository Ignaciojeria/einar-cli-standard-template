package firestore_repository

import (
	"archetype/app/infrastructure/firebasewrapper/firestorewrapper"
	"context"

	"cloud.google.com/go/firestore"
	"go.opentelemetry.io/otel/trace"
)

func RunFirestoreOperation(ctx context.Context, entity interface{}) error {
	_, span := tracer.Start(ctx,
		"RunFirestoreOperation",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	var _ *firestore.CollectionRef = firestorewrapper.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
	//Do something with collection.
	return nil
}
