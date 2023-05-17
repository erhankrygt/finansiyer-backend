package service

import (
	"context"
	"errors"
	"finansiyer"
	mockmongostore "finansiyer/internal/mock/store/mongo"
	mongostore "finansiyer/internal/store/mongo"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

const (
	testCompanyGuid = "3006e084-9952-4083-bdf6-6d67dec9c2d4"
)

func TestHealth(t *testing.T) {
	req := finansiyer.HealthRequest{}
	expectedRes := finansiyer.HealthResponse{
		Data: &finansiyer.HealthData{
			Ping: "pong",
		},
	}

	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	if !reflect.DeepEqual(expectedRes, s.Health(ctx, req)) {
		t.Errorf("expected response and response don't match, expected response: %+v", expectedRes)
	}
}
func Test_CommercialPaperSuccess(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}
	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      true,
		EInvoice:    true,
		TaxNumber:   "0123456789",
	}

	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	dueDate, _ := convertStringToDate("01.02.2024")

	insertDocumentRequest := mongostore.Document{
		Company: mongostore.Company{
			CompanyGuid: testCompanyGuid,
			Status:      true,
			EInvoice:    true,
			TaxNumber:   "0123456789",
		},
		Amount:    1000,
		Currency:  "TRY",
		DueDate:   dueDate,
		BankRate:  BankRate,
		Issuer:    MockUser,
		CreatedAt: getCreatedAt(),
	}

	insertDocumentResponse := &mongostore.Document{
		Company: mongostore.Company{
			CompanyGuid: testCompanyGuid,
			Status:      true,
			EInvoice:    true,
			TaxNumber:   "0123456789",
		},
		Amount:    1000,
		Currency:  "TRY",
		DueDate:   dueDate,
		BankRate:  BankRate,
		Issuer:    MockUser,
		CreatedAt: getCreatedAt(),
	}

	ms.On("InsertDocument", ctx, insertDocumentRequest).Return(insertDocumentResponse, nil)

	insertOperationRequest := mongostore.Operation{
		Document:  insertDocumentRequest,
		Status:    mongostore.Review,
		CreatedAt: getCreatedAt(),
		Issuer:    MockUser,
	}
	ms.On("InsertOperation", ctx, insertOperationRequest).Return(true, nil)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2024",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data: &finansiyer.CommercialPaperData{
			Success: true,
		},
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperDocumentNotValid(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2022",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}

	getCompanyResponse := &mongostore.Company{}

	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, errors.New("occurred an error"))

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: ErrDocumentNotLegal.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperTaxNumberNotValid(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2022",
		TaxNumber:   "123456789",
		CompanyGuid: testCompanyGuid,
	}

	_, err := validTaxNumber(req.TaxNumber)

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: err.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperCurrencyNotValid(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TR",
		Amount:      1000,
		DueDate:     "01.02.2022",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	_, err := validatedCurrency(req.Currency)

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: err.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperDueDateNotValid(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02-2022",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	_, err := convertStringToDate(req.DueDate)

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: err.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperDueDateBeforeNowError(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}

	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      false,
		EInvoice:    false,
		TaxNumber:   "0123456789",
	}
	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2022",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: ErrDocumentNotLegal.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperEInvoiceError(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}

	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      true,
		EInvoice:    false,
		TaxNumber:   "0123456789",
	}
	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2024",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: ErrDocumentNotLegal.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperCompanyStatusError(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}

	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      false,
		EInvoice:    true,
		TaxNumber:   "0123456789",
	}
	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2024",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: ErrDocumentNotLegal.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperTaxNumberError(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}

	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      true,
		EInvoice:    true,
		TaxNumber:   "01234567890",
	}
	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2024",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: ErrDocumentNotLegal.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperInsertDocumentError(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}
	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      true,
		EInvoice:    true,
		TaxNumber:   "0123456789",
	}

	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	dueDate, _ := convertStringToDate("01.02.2024")

	insertDocumentRequest := mongostore.Document{
		Company: mongostore.Company{
			CompanyGuid: testCompanyGuid,
			Status:      true,
			EInvoice:    true,
			TaxNumber:   "0123456789",
		},
		Amount:    1000,
		Currency:  "TRY",
		DueDate:   dueDate,
		BankRate:  BankRate,
		Issuer:    MockUser,
		CreatedAt: getCreatedAt(),
	}

	insertDocumentResponse := &mongostore.Document{}

	err := errors.New("occurred an error")
	ms.On("InsertDocument", ctx, insertDocumentRequest).Return(insertDocumentResponse, err)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2024",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data:   nil,
		Result: err.Error(),
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}

func Test_CommercialPaperInsertOperationError(t *testing.T) {
	l := log.NewLogfmtLogger(os.Stdout)
	ms := mockmongostore.NewStore()
	ctx := context.Background()

	s := NewService(l, ms, "dev")

	getCompanyRequest := mongostore.FindCompanyFilter{
		CompanyGuid: testCompanyGuid,
	}
	getCompanyResponse := &mongostore.Company{
		CompanyGuid: testCompanyGuid,
		Status:      true,
		EInvoice:    true,
		TaxNumber:   "0123456789",
	}

	ms.On("GetCompany", ctx, getCompanyRequest).Return(getCompanyResponse, nil)

	dueDate, _ := convertStringToDate("01.02.2024")

	insertDocumentRequest := mongostore.Document{
		Company: mongostore.Company{
			CompanyGuid: testCompanyGuid,
			Status:      true,
			EInvoice:    true,
			TaxNumber:   "0123456789",
		},
		Amount:    1000,
		Currency:  "TRY",
		DueDate:   dueDate,
		BankRate:  BankRate,
		Issuer:    MockUser,
		CreatedAt: getCreatedAt(),
	}

	insertDocumentResponse := &mongostore.Document{
		Company: mongostore.Company{
			CompanyGuid: testCompanyGuid,
			Status:      true,
			EInvoice:    true,
			TaxNumber:   "0123456789",
		},
		Amount:    1000,
		Currency:  "TRY",
		DueDate:   dueDate,
		BankRate:  BankRate,
		Issuer:    MockUser,
		CreatedAt: getCreatedAt(),
	}

	ms.On("InsertDocument", ctx, insertDocumentRequest).Return(insertDocumentResponse, nil)

	insertOperationRequest := mongostore.Operation{
		Document:  insertDocumentRequest,
		Status:    mongostore.Review,
		CreatedAt: getCreatedAt(),
		Issuer:    MockUser,
	}

	err := errors.New("occurred an error")

	ms.On("InsertOperation", ctx, insertOperationRequest).Return(false, err)

	req := finansiyer.CommercialPaperRequest{
		Currency:    "TRY",
		Amount:      1000,
		DueDate:     "01.02.2024",
		TaxNumber:   "0123456789",
		CompanyGuid: testCompanyGuid,
	}

	expectedResponse := finansiyer.CommercialPaperResponse{
		Data: &finansiyer.CommercialPaperData{
			Success: false,
		},
	}

	response := s.CommercialPaper(ctx, req)
	assert.Equal(t, expectedResponse, response)
}
