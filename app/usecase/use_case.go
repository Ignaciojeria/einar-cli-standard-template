package usecase

import (
	"archetype/app/shared/infrastructure/observability"
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type INewUsecase func(ctx context.Context, input interface{}) (interface{}, error)

func init() {
	ioc.Registry(NewUseCase)
}

func NewUseCase() INewUsecase {
	return func(ctx context.Context, input interface{}) (interface{}, error) {
		_, span := observability.Tracer.Start(ctx, "INewUsecase")
		defer span.End()
		return input, nil
	}
}
