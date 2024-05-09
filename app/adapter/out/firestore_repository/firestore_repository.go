package firestore_repository

import (
	"archetype/app/shared/infrastructure/firebasewrapper/firestorewrapper"
	"archetype/app/shared/infrastructure/observability"
	"context"

	"cloud.google.com/go/firestore"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/trace"
)

type IRunFirestoreOperation func(ctx context.Context, input interface{}) error

func init() {
	ioc.Registry(
		NewRunFirestoreOperation,
		firestorewrapper.NewClientWrapper)
}
func NewRunFirestoreOperation(c *firestorewrapper.ClientWrapper) IRunFirestoreOperation {
	return func(ctx context.Context, input interface{}) error {
		_, span := observability.Tracer.Start(ctx,
			"IRunFirestoreOperation",
			trace.WithSpanKind(trace.SpanKindInternal))
		defer span.End()
		var _ *firestore.CollectionRef = c.Collection("INSERT_YOUR_COLLECTION_CONSTANT_HERE")
		//Do something with collection.
		return nil
	}
}
