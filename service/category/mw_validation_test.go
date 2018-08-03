// +build unit

package category

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/minhkhiemm/example-go/domain"
)

func Test_validationMiddleware_Update(t *testing.T) {
	serviceMock := &ServiceMock{
		UpdateFunc: func(_ context.Context, p *domain.Category) (*domain.Category, error) {
			return p, nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name            string
		args            args
		wantOutput      *domain.Category
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid category",
			args: args{&domain.Category{
				Name: "Curabitur vulputate vestibulum lorem.",
			}},
			wantOutput: &domain.Category{
				Name: "Curabitur vulputate vestibulum lorem.",
			},
		},
		{
			name:            "invalid category by missing name",
			args:            args{&domain.Category{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid category by missing attribute",
			args:            args{&domain.Category{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
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
		CreateFunc: func(_ context.Context, p *domain.Category) error {
			return nil
		},
	}

	defaultCtx := context.Background()
	type args struct {
		p *domain.Category
	}
	tests := []struct {
		name            string
		args            args
		wantErr         bool
		errorStatusCode int
	}{
		{
			name: "valid category",
			args: args{&domain.Category{
				Name: "Curabitur vulputate vestibulum lorem.",
			}},
		},
		{
			name:            "invalid category by missing name",
			args:            args{&domain.Category{}},
			wantErr:         true,
			errorStatusCode: http.StatusBadRequest,
		},
		{
			name:            "invalid category by missing attribute",
			args:            args{&domain.Category{}},
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
		p   *domain.Category
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput *domain.Category
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
		wantOutput []domain.Category
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
		p   *domain.Category
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

func TestValidationMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want func(Service) Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidationMiddleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidationMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
