package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"wallet"
)

// Endpoints represents service endpoints
type Endpoints struct {
	HealthEndpoint endpoint.Endpoint
}

// MakeEndpoints makes and returns endpoints
func MakeEndpoints(s wallet.Service) Endpoints {
	return Endpoints{
		HealthEndpoint: MakeHealthEndpoint(s),
	}
}

// MakeHealthEndpoint makes and returns health endpoint
func MakeHealthEndpoint(s wallet.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*wallet.HealthRequest)

		res := s.Health(ctx, *req)

		return res, nil
	}
}
