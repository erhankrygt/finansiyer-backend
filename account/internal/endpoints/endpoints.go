package endpoints

import (
	"context"
	"github.com/erhankrygt/finansiyer-backend/account"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints represents service endpoints
type Endpoints struct {
	HealthEndpoint endpoint.Endpoint
}

// MakeEndpoints makes and returns endpoints
func MakeEndpoints(s account.Service) Endpoints {
	return Endpoints{
		HealthEndpoint: MakeHealthEndpoint(s),
	}
}

// MakeHealthEndpoint makes and returns health endpoint
func MakeHealthEndpoint(s account.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*account.HealthRequest)

		res := s.Health(ctx, *req)

		return res, nil
	}
}
