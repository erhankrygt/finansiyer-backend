package service

import (
	"context"
	"github.com/go-kit/log"
	"search"
	envvars "search/configs/env-vars"
	mongostore "search/internal/store/mongo"
)

// compile-time proofs of service interface implementation
var _ search.Service = (*RestService)(nil)

// RestService represents service
type RestService struct {
	l   log.Logger
	ms  mongostore.Store
	env envvars.EnvVars
}

// NewService creates and returns service
func NewService(l log.Logger, ms mongostore.Store, env envvars.EnvVars) search.Service {
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
func (s *RestService) Health(_ context.Context, _ search.HealthRequest) search.HealthResponse {
	return search.HealthResponse{
		Data: &search.HealthData{
			Ping: "pong",
		},
	}
}
