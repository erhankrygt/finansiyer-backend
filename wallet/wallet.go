package wallet

import "context"

// Service defines behaviors of sample service
type Service interface {
	Health(context.Context, HealthRequest) HealthResponse
	CreateBank(context.Context, CreateBankRequest) CreateBankResponse
	GetBanks(context.Context, BankRequest) BankResponse
	CreateBankAccount(context.Context, CreateBankAccountRequest) CreateBankAccountResponse
}

type Header struct {
	UserID    string `json:"userID"`
	UserToken string `json:"userToken"`
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
	BankRequest struct {
		Header
	}

	BankResponse struct {
		Data   *BankData `json:"data"`
		Result *ApiError `json:"result,omitempty"`
	}

	BankData struct {
		Banks []Bank `json:"banks"`
	}

	Bank struct {
		Title   string `json:"title"`
		WebSite string `json:"webSite"`
	}
)

// CreateBankRequest and CreateBankResponse represents create bank request and response
type (
	CreateBankRequest struct {
		Header
		Title   string `json:"title"`
		WebSite string `json:"webSite"`
	}

	CreateBankResponse struct {
		Data   *CreateBankData `json:"data"`
		Result *ApiError       `json:"result,omitempty"`
	}

	CreateBankData struct {
		Success bool `json:"success"`
	}
)

// CreateBankAccountRequest and CreateBankAccountResponse represents create bank account request and response
type (
	CreateBankAccountRequest struct {
		Header
		BankID        string `json:"bankID"`
		IBAN          string `json:"iban"`
		AccountNumber string `json:"accountNumber"`
	}

	CreateBankAccountResponse struct {
		Data   *CreateBankAccountData `json:"data"`
		Result *ApiError              `json:"result,omitempty"`
	}

	CreateBankAccountData struct {
		Success bool `json:"success"`
	}
)

type ApiError struct {
	Code    int
	Message *error
}
