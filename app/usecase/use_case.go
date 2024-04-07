package usecase

import (
	"context"

	ioc "github.com/Ignaciojeria/einar-ioc"
)

type IUseCase interface {
	Execute(ctx context.Context, domain interface{}) (interface{}, error)
}

type useCase struct {
}

func init() {
	ioc.Registry(newUseCase)
}
func newUseCase() IUseCase {
	return useCase{}
}

func (u useCase) Execute(ctx context.Context, domain interface{}) (interface{}, error) {
	// business logic here
	return nil, nil
}

func Instance() IUseCase {
	return ioc.Get[IUseCase](newUseCase)
}
