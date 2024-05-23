package postgresql_repository

import (
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/infrastructure/postgresql"
	"archetype/app/shared/logging"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type RunPostgreSQLOperation func(ctx context.Context, input interface{}) error

func init() {
	ioc.Registry(
		NewRunPostgreSQLOperation,
		postgresql.NewConnection,
		logging.NewLogger)
}
func NewRunPostgreSQLOperation(connection *gorm.DB, logger logging.Logger) RunPostgreSQLOperation {
	return func(ctx context.Context, input interface{}) error {
		_, span := observability.Tracer.Start(ctx,
			"RunPostgreSQLOperation",
			trace.WithSpanKind(trace.SpanKindInternal))
		defer span.End()
		//PUT YOUR FIRESTORE OPERATION HERE
		return nil
	}
}
