package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/erhankrygt/finansiyer-backend/account"
	envvars "github.com/erhankrygt/finansiyer-backend/account/configs/env-vars"
	mongostore "github.com/erhankrygt/finansiyer-backend/account/internal/store/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	env envvars.EnvVars
}

// NewService creates and returns service
func NewService(l log.Logger, ms mongostore.Store, env envvars.EnvVars) account.Service {
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
	verify, _ := verifyPassword(req.Password, req.ConfirmPassword)
	if verify == false {
		return account.RegisterResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: &ErrPasswordsDoNotMatch,
			},
		}
	}

	user, err := s.ms.GetUser(ctx, req.PhoneNumber)
	if err != nil {
		// TODO: Log with Sentry
	}

	if user != nil {
		return account.RegisterResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: &ErrAlreadyRegistered,
			},
		}
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		// TODO: Log with Sentry
		return account.RegisterResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: &err,
			},
		}
	}

	claims := jwt.MapClaims{
		"phonenumber": req.PhoneNumber,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte("gizli_anahtar"))

	u := mongostore.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		Password:    hashedPassword,
		CreatedAt:   getCreatedAt(),
		Token:       token,
		IsActive:    true,
		IsDeleted:   false,
	}

	_, err = s.ms.InsertUser(ctx, u)
	if err != nil {
		// TODO: Log with Sentry

		return account.RegisterResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: &err,
			},
		}
	}

	return account.RegisterResponse{
		Data: &account.RegisterData{
			IsSuccessful: true,
		},
	}
}

// Login returns login
// swagger:operation GET /account/login loginRequest
// ---
// summary: Login
// description: Returns response of login result
// responses:
//
//	  200:
//		  $ref: "#/responses/loginResponse"
func (s *RestService) Login(ctx context.Context, req account.LoginRequest) account.LoginResponse {
	user, err := s.ms.GetUser(ctx, req.PhoneNumber)
	if err != nil {
		// TODO: Log with Sentry
		return account.LoginResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: &err,
			},
		}
	}

	err = comparePasswords(user.Password, req.Password)
	if err != nil {
		return account.LoginResponse{
			Data: nil,
			Result: &account.ApiError{
				Code:    http.StatusBadRequest,
				Message: &ErrLoginInformationInCorrect,
			},
		}
	}

	return account.LoginResponse{
		Data: &account.LoginData{
			IsSuccessful: true,
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

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func comparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
