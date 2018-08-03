// +build unit

package book

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	domain "github.com/minhkhiemm/example-go/domain"
)

func Test_validationMiddleware_Update(t *testing.T) {
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Book) (*domain.Book, error) {
			return p, nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Book
	}
	tests := []struct {
		name            string
		category_id     domain.UUID
		description     string
		args            args
		wantOutput      *domain.Book
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
				CategoryID:  domain.MustGetUUIDFromString("4fb9c805-ae37-4eec-8e61-3e4d39686a87"),
			}},
			wantOutput: &domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
				CategoryID:  domain.MustGetUUIDFromString("4fb9c805-ae37-4eec-8e61-3e4d39686a87"),
			},
		},
		{
			name: "invalid book by missing name",
			args: args{&domain.Book{
				Description: "Curabitur vulputate vestibulum lorem.",
				CategoryID:  domain.MustGetUUIDFromString("4fb9c805-ae37-4eec-8e61-3e4d39686a87"),
			}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid book by missing attribute",
			args:            args{&domain.Book{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by missing description",
			args: args{&domain.Book{
				Name:       "Curabitur vulputate vestibulum lorem.",
				CategoryID: domain.MustGetUUIDFromString("4fb9c805-ae37-4eec-8e61-3e4d39686a87"),
			}},
			wantErr:         ErrDescriptionIsRequired,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by invalid description",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				CategoryID:  domain.MustGetUUIDFromString("4fb9c805-ae37-4eec-8e61-3e4d39686a87"),
				Description: "six",
			}},
			wantErr:         ErrDescriptionIsInvalid,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid book by missing category_id",
			args: args{&domain.Book{
				Name:        "Curabitur vulputate vestibulum lorem.",
				Description: "Curabitur vulputate vestibulum lorem.",
			}},
			wantErr true,
			errorStatusCode: http.StatusBadRequest,
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: serviceMock,
			}
			gotOutput, err := mw.Update(defaultCtx, tt.args.p)
			if err != nil {
				if tt.wantErr == false {
					t.Errorf("validationMiddleware.Update() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				status, ok := err.(interface{ StatusCode() int })
				if !ok {
					t.Errorf("validationMiddleware.Update() error %v doesn't implement StatusCode()", err)
				}
				if tt.errorStatusCode != status.StatusCode() {
					t.Errorf("validationMiddleware.Update() status = %v, want status code %v", status.StatusCode(), tt.errorStatusCode)
					return
				}

				return
			}

			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("ValidationMiddleware.Update() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

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
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid book",
			args: args{&domain.Book{
				Name: "Curabitur vulputate vestibulum lorem.",
			}},
		},
		{
			name:            "invalid book by missing name",
			args:            args{&domain.Book{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid book by missing attribute",
			args:            args{&domain.Book{}},
			wantErr:         true,
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
				if tt.wantErr == false {
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

func Test_validationMiddleware_Find(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.Book
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput *domain.Book
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			gotOutput, err := mw.Find(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("validationMiddleware.Find() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_validationMiddleware_FindAll(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput []domain.Book
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			gotOutput, err := mw.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("validationMiddleware.FindAll() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func Test_validationMiddleware_Delete(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		ctx context.Context
		p   *domain.Book
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: tt.fields.Service,
			}
			if err := mw.Delete(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("validationMiddleware.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
