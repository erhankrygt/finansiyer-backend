package service

import (
	"context"
	"github.com/go-kit/log"
	"net/http"
	"wallet"
	envvars "wallet/configs/env-vars"
	mongostore "wallet/internal/store/mongo"
)

// compile-time proofs of service interface implementation
var _ wallet.Service = (*RestService)(nil)

// RestService represents service
type RestService struct {
	l   log.Logger
	ms  mongostore.Store
	env envvars.EnvVars
}

// NewService creates and returns service
func NewService(l log.Logger, ms mongostore.Store, env envvars.EnvVars) wallet.Service {
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
func (s *RestService) Health(_ context.Context, _ wallet.HealthRequest) wallet.HealthResponse {
	return wallet.HealthResponse{
		Data: &wallet.HealthData{
			Ping: "pong",
		},
	}
}

// CreateBank returns create bank
// swagger:operation POST /create-bank createBankRequest
// ---
// summary: CreateBank
// description: Returns response of create bank result
// responses:
//
//	  200:
//		  $ref: "#/responses/createBankResponse"
func (s *RestService) CreateBank(ctx context.Context, req wallet.CreateBankRequest) wallet.CreateBankResponse {

	res := wallet.CreateBankResponse{}

	bank := mongostore.Bank{
		Title:   req.Title,
		WebSite: req.WebSite,
		User: mongostore.User{
			UserID: req.UserID,
		},
	}

	success, err := s.ms.CreateBank(ctx, bank)
	if err != nil {
		res.Result = &wallet.ApiError{
			Code:    http.StatusBadRequest,
			Message: &err,
		}
	}

	res.Data = &wallet.CreateBankData{
		Success: success,
	}

	return res
}

// GetBanks returns banks
// swagger:operation GET /banks bankRequest
// ---
// summary: GetBanks
// description: Returns response of bank result
// responses:
//
//	  200:
//		  $ref: "#/responses/bankResponse"
func (s *RestService) GetBanks(ctx context.Context, req wallet.BankRequest) wallet.BankResponse {

	res := wallet.BankResponse{}

	user := mongostore.User{
		UserID: req.UserID,
	}

	bankList, err := s.ms.GetBanks(ctx, user)
	if err != nil {
		res.Result = &wallet.ApiError{
			Code:    http.StatusBadRequest,
			Message: &err,
		}
	}

	var banks []wallet.Bank

	for _, b := range bankList {
		banks = append(banks, wallet.Bank{
			Title:   b.Title,
			WebSite: b.WebSite,
		})
	}

	res.Data = &wallet.BankData{
		Banks: banks,
	}

	return res
}

func (s *RestService) CreateBankAccount(ctx context.Context, req wallet.CreateBankAccountRequest) wallet.CreateBankAccountResponse {
	//TODO implement me
	panic("implement me")
}
