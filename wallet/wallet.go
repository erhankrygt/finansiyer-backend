package wallet

import "context"

// Service defines behaviors of sample service
type Service interface {
	Health(context.Context, HealthRequest) HealthResponse
	GetBanks(context.Context, BankRequest) BankResponse
}

// Request defines behaviors of request
type Request interface{}

// Response defines behaviors of response
type Response interface{}

// compile-time proofs of request interface implementation
var (
	_ Request = (*HealthRequest)(nil)
)

// compile-time proofs of response interface implementation
var (
	_ Response = (*HealthResponse)(nil)
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

// BankRequest and BankResponse represents bank request and response
type (
	BankRequest struct{}

	BankResponse struct {
		Data   *BankData `json:"data"`
		Result *ApiError `json:"result,omitempty"`
	}

	BankData struct {
		Title   string `json:"title"`
		WebSite string `json:"webSite"`
	}
)

type ApiError struct {
	Code    int
	Message *error
}
