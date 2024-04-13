package usecase

import (
	"archetype/app/infrastructure/observability"
	"context"
)

func NewUseCase(ctx context.Context, domain interface{}) (interface{}, error) {
	_, span := observability.Tracer.Start(ctx, "NewUseCase")
	defer span.End()
	return "Unimplemented", nil
}
