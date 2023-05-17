package mockmongostore

import (
	mongostore "github.com/erhankrygt/finansiyer-backend/account/internal/store/mongo"
	"github.com/stretchr/testify/mock"
)

// compile-time proof of mongo store interface implementation
var _ mongostore.Store = (*Store)(nil)

// Store represents mock mongo store
type Store struct {
	mock.Mock
}

// NewStore returns mock mongo store
func NewStore() *Store {
	return &Store{}
}

// Close mocks close method
func (s Store) Close() error {
	args := s.Called()
	return args.Error(0)
}
