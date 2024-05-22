package firestore_repository

import (
	"archetype/app/shared/infrastructure/firebaseapp/firestoreclient"
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/logging"
	"context"

	"cloud.google.com/go/firestore"
	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/trace"
)

type IRunFirestoreOperation func(ctx context.Context, input interface{}) error

func init() {
	ioc.Registry(
		NewRunFirestoreOperation,
		firestoreclient.NewClient,
		logging.NewLogger)
}
func NewRunFirestoreOperation(c *firestore.Client, logger logging.Logger) IRunFirestoreOperation {
	return func(ctx context.Context, input interface{}) error {
		_, span := observability.Tracer.Start(ctx,
			"IRunFirestoreOperation",
			trace.WithSpanKind(trace.SpanKindInternal))
		defer span.End()
		//PUT YOUR FIRESTORE OPERATION HERE
		return nil
	}
}
