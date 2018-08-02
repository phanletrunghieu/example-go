package book

import (
	"context"
	"net/http"
	"testing"

	"github.com/phanletrunghieu/example-go/domain"
)

func Test_validationMiddleware_Create(t *testing.T) {
	serviceMock := &ServiceMock{
		CreateFunc: func(_ context.Context, p *domain.Book) error {
			return nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur",
			}},
		},
		{
			name: "name book empty",
			args: args{&domain.Book{
				Description: "Curabitur",
			}},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of name <= 5 (1)",
			args: args{&domain.Book{
				Name:        "Curab",
				Description: "Curabitur",
			}},
			wantErr:         ErrNameLenght,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of name <= 5 (2)",
			args: args{&domain.Book{
				Name:        "Cur",
				Description: "Curabitur",
			}},
			wantErr:         ErrNameLenght,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "description book empty",
			args: args{&domain.Book{
				Name: "Curabitur",
			}},
			wantErr:         ErrDescriptionLength,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of description <= 5 (1)",
			args: args{&domain.Book{
				Name:        "Curabitur",
				Description: "Curab",
			}},
			wantErr:         ErrDescriptionLength,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of description <= 5 (2)",
			args: args{&domain.Book{
				Name:        "Curabitur",
				Description: "Cur",
			}},
			wantErr:         ErrDescriptionLength,
			errorStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			err := mw.Create(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Create() error %v doesn't implement StatusCode()", err)
				}
				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Create() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}
		})
	}
}

func Test_validationMiddleware_Update(t *testing.T) {
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Book) (*domain.Book, error) {
			return nil, nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name            string
		args            args
		wantErr         error
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur",
			}},
		},
		{
			name: "name book empty",
			args: args{&domain.Book{
				Description: "Curabitur",
			}},
			wantErr:         ErrNameIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of name <= 5 (1)",
			args: args{&domain.Book{
				Name:        "Curab",
				Description: "Curabitur",
			}},
			wantErr:         ErrNameLenght,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of name <= 5 (2)",
			args: args{&domain.Book{
				Name:        "Cur",
				Description: "Curabitur",
			}},
			wantErr:         ErrNameLenght,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "description book empty",
			args: args{&domain.Book{
				Name: "Curabitur",
			}},
			wantErr:         ErrDescriptionLength,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of description <= 5 (1)",
			args: args{&domain.Book{
				Name:        "Curabitur",
				Description: "Curab",
			}},
			wantErr:         ErrDescriptionLength,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "length of description <= 5 (2)",
			args: args{&domain.Book{
				Name:        "Curabitur",
				Description: "Cur",
			}},
			wantErr:         ErrDescriptionLength,
			errorStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			_, err := mw.Update(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr != err {
					t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Create() error %v doesn't implement StatusCode()", err)
				}
				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Create() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	fakeID := domain.MustGetUUIDFromString("1698bbd6-e0c8-4957-a5a9-8c536970994b")

	serviceMock := &ServiceMock{
		DeleteFunc: func(_ context.Context, p *domain.Book) error {
			return nil
		},
	}

	type args struct {
		book *domain.Book
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "success",
			args: args{&domain.Book{
				Model: domain.Model{ID: fakeID},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			if err := mw.Delete(context.Background(), tt.args.book); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
