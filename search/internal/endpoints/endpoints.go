package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"search"
)

// Endpoints represents service endpoints
type Endpoints struct {
	HealthEndpoint endpoint.Endpoint
	BlogEndpoint   endpoint.Endpoint
}

// MakeEndpoints makes and returns endpoints
func MakeEndpoints(s search.Service) Endpoints {
	return Endpoints{
		HealthEndpoint: MakeHealthEndpoint(s),
		BlogEndpoint:   MakeBlogEndpoint(s),
	}
}

// MakeHealthEndpoint makes and returns health endpoint
func MakeHealthEndpoint(s search.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*search.HealthRequest)

		res := s.Health(ctx, *req)

		return res, nil
	}
}

// MakeBlogEndpoint makes and returns blog endpoint
func MakeBlogEndpoint(s search.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*search.HealthRequest)

		res := s.Health(ctx, *req)

		return res, nil
	}
}
