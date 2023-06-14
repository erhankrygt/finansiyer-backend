package service

import (
	"context"
	"github.com/go-kit/log"
	"wallet"
	envvars "wallet/configs/env-vars"
	mongostore "wallet/internal/store/mongo"
)

// compile-time proofs of service interface implementation
var _ wallet.Service = (*RestService)(nil)

// RestService represents service
type RestService struct {
	l   log.Logger
	ms  mongostore.Store
	env envvars.EnvVars
}

// NewService creates and returns service
func NewService(l log.Logger, ms mongostore.Store, env envvars.EnvVars) wallet.Service {
	return &RestService{
		l:   l,
		ms:  ms,
		env: env,
	}
}

// Health returns health
// swagger:operation GET /health healthRequest
// ---
// summary: Health
// description: Returns response of health result
// responses:
//
//	  200:
//		  $ref: "#/responses/healthResponse"
func (s *RestService) Health(_ context.Context, _ wallet.HealthRequest) wallet.HealthResponse {
	return wallet.HealthResponse{
		Data: &wallet.HealthData{
			Ping: "pong",
		},
	}
}

// GetBanks returns banks
// swagger:operation GET /banks bankRequest
// ---
// summary: GetBanks
// description: Returns response of bank result
// responses:
//
//	  200:
//		  $ref: "#/responses/bankResponse"
func (s *RestService) GetBanks(ctx context.Context, request wallet.BankRequest) wallet.BankResponse {
	//TODO implement me
	panic("implement me")
}
