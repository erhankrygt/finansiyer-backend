package account

import (
	"context"
)

// Service defines behaviors of sample service
type Service interface {
	Health(context.Context, HealthRequest) HealthResponse
	Register(context.Context, RegisterRequest) RegisterResponse
	Login(context.Context, LoginRequest) LoginResponse
}

// Request defines behaviors of request
type Request interface{}

// Response defines behaviors of response
type Response interface{}

// compile-time proofs of request interface implementation
var (
	_ Request = (*HealthRequest)(nil)
	_ Request = (*RegisterRequest)(nil)
	_ Request = (*LoginRequest)(nil)
)

// compile-time proofs of response interface implementation
var (
	_ Response = (*HealthResponse)(nil)
	_ Response = (*RegisterResponse)(nil)
	_ Response = (*LoginResponse)(nil)
)

// HealthRequest and HealthResponse represents health request and response
type (
	HealthRequest struct{}

	HealthResponse struct {
		Data   *HealthData `json:"data"`
		Result *ApiError   `json:"result,omitempty"`
	}

	HealthData struct {
		Ping string `json:"ping"`
	}
)

// RegisterRequest and RegisterResponse represents register request and response
type (
	RegisterRequest struct {
		FirstName       string `json:"firstName" validate:"required"`
		LastName        string `json:"lastName" validate:"required"`
		Email           string `json:"email" validate:"required"`
		PhoneNumber     string `json:"phoneNumber" validate:"required"`
		Password        string `json:"password" validate:"required"`
		ConfirmPassword string `json:"confirmPassword" validate:"required"`
	}

	RegisterResponse struct {
		Data   *RegisterData `json:"data"`
		Result *ApiError     `json:"result,omitempty"`
	}

	RegisterData struct {
		IsSuccessful bool `json:"isSuccessful"`
	}
)

// LoginRequest and LoginResponse represents login request and response
type (
	LoginRequest struct {
		PhoneNumber string `json:"phoneNumber" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Data   *LoginData `json:"data"`
		Result *ApiError  `json:"result,omitempty"`
	}

	LoginData struct {
		IsSuccessful bool `json:"isSuccessful"`
	}
)

type ApiError struct {
	Code    int
	Message *error
}
