package account

import (
	"context"
)

// Service defines behaviors of sample service
type Service interface {
	Health(context.Context, HealthRequest) HealthResponse
	Register(context.Context, RegisterRequest) RegisterResponse
}

// Request defines behaviors of request
type Request interface{}

// Response defines behaviors of response
type Response interface{}

// compile-time proofs of request interface implementation
var (
	_ Request = (*HealthRequest)(nil)
	_ Request = (*RegisterRequest)(nil)
)

// compile-time proofs of response interface implementation
var (
	_ Response = (*HealthResponse)(nil)
	_ Response = (*RegisterResponse)(nil)
)

// HealthRequest and HealthResponse represents health request and response
type (
	HealthRequest struct{}

	HealthResponse struct {
		Data   *HealthData `json:"data"`
		Result string      `json:"result,omitempty"`
	}

	HealthData struct {
		Ping string `json:"ping"`
	}
)

// RegisterRequest and RegisterResponse represents register request and response
type (
	RegisterRequest struct {
		FirstName       string `json:"firstName"`
		LastName        string `json:"lastName"`
		PhoneNumber     string `json:"phoneNumber"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	RegisterResponse struct {
		Data   *RegisterData `json:"data"`
		Result string        `json:"result,omitempty"`
	}

	RegisterData struct {
		IsSuccessful string `json:"isSuccessful"`
	}
)
