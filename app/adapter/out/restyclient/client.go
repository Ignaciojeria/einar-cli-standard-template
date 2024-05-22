package restyclient

import (
	"archetype/app/shared/infrastructure/newresty"
	"archetype/app/shared/infrastructure/observability"
	"archetype/app/shared/logging"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel/trace"
)

type Client func(ctx context.Context, input interface{}) (interface{}, error)

func init() {
	ioc.Registry(NewClient, newresty.NewClient, logging.NewLogger)
}
func NewClient(cli *resty.Client, logger logging.Logger) Client {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		_, span := observability.Tracer.Start(ctx,
			"Client",
			trace.WithSpanKind(trace.SpanKindInternal))
		defer span.End()
		return nil, nil
	}
}
