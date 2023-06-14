package service

import (
	"context"
	"github.com/erhankrygt/finansiyer-backend/account"
	envvars "github.com/erhankrygt/finansiyer-backend/account/configs/env-vars"
	mockmongostore "github.com/erhankrygt/finansiyer-backend/account/internal/mock/store/mongo"
	"github.com/go-kit/log"
	"os"
	"reflect"
	"testing"
)

const (
	testCompanyGuid = "3006e084-9952-4083-bdf6-6d67dec9c2d4"
)

func TestHealth(t *testing.T) {
	req := account.HealthRequest{}
	expectedRes := account.HealthResponse{
		Data: &account.HealthData{
			Ping: "pong",
		},
	}

	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, envvars.EnvVars{})

	if !reflect.DeepEqual(expectedRes, s.Health(ctx, req)) {
		t.Errorf("expected response and response don't match, expected response: %+v", expectedRes)
	}
}
