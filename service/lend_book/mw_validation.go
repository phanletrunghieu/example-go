package lendbook

import (
	"context"

	"github.com/phanletrunghieu/example-go/domain"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, lendBook *domain.LendBook) (err error) {
	return mw.Service.Create(ctx, lendBook)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.LendBook, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, lendBook *domain.LendBook) (*domain.LendBook, error) {
	return mw.Service.Find(ctx, lendBook)
}

func (mw validationMiddleware) Update(ctx context.Context, lendBook *domain.LendBook) (*domain.LendBook, error) {
	return mw.Service.Update(ctx, lendBook)
}
func (mw validationMiddleware) Delete(ctx context.Context, lendBook *domain.LendBook) error {
	return mw.Service.Delete(ctx, lendBook)
}
