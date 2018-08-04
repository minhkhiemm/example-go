package category

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/service"
	categoryService "github.com/minhkhiemm/example-go/service/category"
)

func TestMakeCreateEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			CreateFunc: func(_ context.Context, p *domain.Category) error {
				return nil
			},
		},
	}

	type args struct {
		req CreateRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "update category endpoint success",
			args: args{
				CreateRequest{
					CreateData{
						Name: "category 1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeCreateEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeFindEndPoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			DeleteFunc: func(_ context.Context, p *domain.Category) error {
				return nil
			},
		},
	}

	type args struct {
		req DeleteRequest
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeDeleteEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeFindAllEndpoint(t *testing.T) {
	type args struct {
		s service.Service
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeFindAllEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFindAllEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeUpdateEndpoint(t *testing.T) {
	mock := service.Service{
		CategoryService: &categoryService.ServiceMock{
			UpdateFunc: func(_ context.Context, p *domain.Category) (*domain.Category, error) {
				return p, nil
			},
		},
	}

	type args struct {
		req UpdateRequest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "update category endpoint success",
			args: args{
				UpdateRequest{
					UpdateData{
						ID:   domain.MustGetUUIDFromString("415179ad-8067-4138-9b0d-41e0c68d4376"),
						Name: "category 1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFunc := MakeUpdateEndpoint(mock)
			_, err := gotFunc(context.Background(), tt.args.req)
			// check no error
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestMakeDeleteEndpoint(t *testing.T) {
	type args struct {
		s service.Service
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeDeleteEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeDeleteEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
