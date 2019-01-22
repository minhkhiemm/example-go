package book

import (
	"reflect"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/minhkhiemm/example-go/domain"
	"github.com/minhkhiemm/example-go/service"
)

func TestCreateResponse_StatusCode(t *testing.T) {
	type fields struct {
		Book domain.Book
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateResponse{
				Book: tt.fields.Book,
			}
			if got := c.StatusCode(); got != tt.want {
				t.Errorf("CreateResponse.StatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeCreateEndpoint(t *testing.T) {
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
			if got := MakeCreateEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeCreateEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeFindEndPoint(t *testing.T) {
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
			if got := MakeFindEndPoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeFindEndPoint() = %v, want %v", got, tt.want)
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
			if got := MakeUpdateEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeUpdateEndpoint() = %v, want %v", got, tt.want)
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
