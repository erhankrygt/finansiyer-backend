package service

import (
	"context"
	"github.com/erhankrygt/finansiyer-backend/account"
	mongostore "github.com/erhankrygt/finansiyer-backend/account/internal/store/mongo"

	"github.com/go-kit/log"
)

// compile-time proofs of service interface implementation
var _ account.Service = (*RestService)(nil)

// RestService represents service
type RestService struct {
	l   log.Logger
	ms  mongostore.Store
	env string
}

// NewService creates and returns service
func NewService(l log.Logger, ms mongostore.Store, env string) account.Service {
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
func (s *RestService) Health(_ context.Context, _ account.HealthRequest) account.HealthResponse {
	return account.HealthResponse{
		Data: &account.HealthData{
			Ping: "pong",
		},
	}
}

func (s *RestService) Register(ctx context.Context, req account.RegisterRequest) account.RegisterResponse {
	//TODO implement me
	panic("implement me")
}