package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"wallet"
)

// Endpoints represents service endpoints
type Endpoints struct {
	HealthEndpoint            endpoint.Endpoint
	CreateBankEndpoint        endpoint.Endpoint
	BanksEndpoint             endpoint.Endpoint
	CreateBankAccountEndpoint endpoint.Endpoint
}

// MakeEndpoints makes and returns endpoints
func MakeEndpoints(s wallet.Service) Endpoints {
	return Endpoints{
		HealthEndpoint:            MakeHealthEndpoint(s),
		BanksEndpoint:             MakeBanksEndpoint(s),
		CreateBankEndpoint:        MakeCreateBankEndpoint(s),
		CreateBankAccountEndpoint: MakeCreateBankAccountEndpoint(s),
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

// MakeBanksEndpoint makes and returns bank endpoint
func MakeBanksEndpoint(s wallet.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*wallet.BankRequest)

		res := s.GetBanks(ctx, *req)

		return res, nil
	}
}

// MakeCreateBankEndpoint makes and returns create bank endpoint
func MakeCreateBankEndpoint(s wallet.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*wallet.BankRequest)

		res := s.GetBanks(ctx, *req)

		return res, nil
	}
}

// MakeCreateBankAccountEndpoint makes and returns create bank account endpoint
func MakeCreateBankAccountEndpoint(s wallet.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*wallet.BankRequest)

		res := s.GetBanks(ctx, *req)

		return res, nil
	}
}
