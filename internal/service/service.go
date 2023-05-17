package service

import (
	"context"
	"finansiyer"
	mongostore "finansiyer/internal/store/mongo"
	"github.com/go-kit/log"
)

// compile-time proofs of service interface implementation
var _ finansiyer.Service = (*RestService)(nil)

// RestService represents service
type RestService struct {
	l   log.Logger
	ms  mongostore.Store
	env string
}

// NewService creates and returns service
func NewService(l log.Logger, ms mongostore.Store, env string) finansiyer.Service {
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
func (s *RestService) Health(_ context.Context, _ finansiyer.HealthRequest) finansiyer.HealthResponse {
	return finansiyer.HealthResponse{
		Data: &finansiyer.HealthData{
			Ping: "pong",
		},
	}
}
