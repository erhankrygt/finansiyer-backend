package service

import (
	"context"
	"github.com/erhankrygt/finansiyer-backend/account"
	mongostore "github.com/erhankrygt/finansiyer-backend/account/internal/store/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

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

// Register returns register
// swagger:operation POST /account/register registerRequest
// ---
// summary: Register
// description: Returns response of register result
// responses:
//
//	  200:
//		  $ref: "#/responses/registerResponse"
func (s *RestService) Register(ctx context.Context, req account.RegisterRequest) account.RegisterResponse {
	v, err := verifyPassword(req.Password, req.ConfirmPassword)
	if err != nil {
		// TODO: Log with Sentry

		return account.RegisterResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		}
	}

	if v {
		u := mongostore.User{
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			PhoneNumber: req.PhoneNumber,
			Email:       req.Email,
			Password:    req.Password,
			CreatedAt:   getCreatedAt(),
		}

		v, err = s.ms.InsertUser(ctx, u)
		if err != nil {
			// TODO: Log with Sentry

			return account.RegisterResponse{
				Data: nil,
				Result: &account.ApiError{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			}
		}
	}

	return account.RegisterResponse{
		Data: &account.RegisterData{
			IsSuccessful: v,
		},
	}
}

func getCreatedAt() primitive.DateTime {
	now := time.Now()
	d := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.UTC)
	r := primitive.NewDateTimeFromTime(d)

	return r
}

func verifyPassword(password, confirmPassword string) (bool, error) {
	if len(password) < PasswordMinimumLength {
		return false, ErrPasswordMustBeMinLen
	}

	if password != confirmPassword {
		return false, ErrPasswordsDoNotMatch
	}

	return true, nil
}
