package firestore_repository

import (
	"archetype/app/shared/infrastructure/firebaseapp/firestoreclient"
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
		firestoreclient.NewClient)
}
func NewRunFirestoreOperation(c *firestore.Client) IRunFirestoreOperation {
	return func(ctx context.Context, input interface{}) error {
		_, span := observability.Tracer.Start(ctx,
			"IRunFirestoreOperation",
			trace.WithSpanKind(trace.SpanKindInternal))
		defer span.End()
		//PUT YOUR FIRESTORE OPERATION HERE
		return nil
	}
}
