package endpoints

import (
	"context"
	"github.com/erhankrygt/finansiyer-backend/account"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints represents service endpoints
type Endpoints struct {
	HealthEndpoint   endpoint.Endpoint
	RegisterEndpoint endpoint.Endpoint
	LoginEndpoint    endpoint.Endpoint
}

// MakeEndpoints makes and returns endpoints
func MakeEndpoints(s account.Service) Endpoints {
	return Endpoints{
		HealthEndpoint:   MakeHealthEndpoint(s),
		RegisterEndpoint: MakeRegisterEndpoint(s),
		LoginEndpoint:    MakeLoginEndpoint(s),
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

// MakeRegisterEndpoint makes and returns register endpoint
func MakeRegisterEndpoint(s account.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*account.RegisterRequest)

		res := s.Register(ctx, *req)

		return res, nil
	}
}

// MakeLoginEndpoint makes and returns login endpoint
func MakeLoginEndpoint(s account.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*account.LoginRequest)

		res := s.Login(ctx, *req)

		return res, nil
	}
}
