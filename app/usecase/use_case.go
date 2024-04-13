package usecase

import (
	"context"
)

func NewUseCase(ctx context.Context, domain interface{}) (interface{}, error) {
	_, span := tracer.Start(ctx, "NewUseCase")
	defer span.End()
	return "Unimplemented", nil
}
