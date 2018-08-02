package lendbook

import (
	"context"
	"testing"

	"github.com/phanletrunghieu/example-go/domain"
)

func Test_validationMiddleware_Create(t *testing.T) {
	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.LendBook) error {
			return nil
		},
	}

	type args struct {
		lendBook *domain.LendBook
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "success",
			args: args{&domain.LendBook{
				User_ID: fakeID,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			if err := mw.Create(context.Background(), tt.args.lendBook); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validationMiddleware_Update(t *testing.T) {
	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.LendBook) (*domain.LendBook, error) {
			return nil, nil
		},
	}

	type args struct {
		lendBook *domain.LendBook
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "success",
			args: args{&domain.LendBook{
				User_ID: fakeID,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			if _, err := mw.Update(context.Background(), tt.args.lendBook); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	serviceMock := &ServiceMock{
		DeleteFunc: func(_ context.Context, p *domain.LendBook) error {
			return nil
		},
	}

	type args struct {
		lendBook *domain.LendBook
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "success",
			args: args{&domain.LendBook{
				User_ID: fakeID,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			if err := mw.Delete(context.Background(), tt.args.lendBook); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
