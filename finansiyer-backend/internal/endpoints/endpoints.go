package endpoints

import (
	"context"
	"finansiyer"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints represents service endpoints
type Endpoints struct {
	HealthEndpoint endpoint.Endpoint
}

// MakeEndpoints makes and returns endpoints
func MakeEndpoints(s finansiyer.Service) Endpoints {
	return Endpoints{
		HealthEndpoint: MakeHealthEndpoint(s),
	}
}

// MakeHealthEndpoint makes and returns health endpoint
func MakeHealthEndpoint(s finansiyer.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*finansiyer.HealthRequest)

		res := s.Health(ctx, *req)

		return res, nil
	}
}
