package mockmongostore

import (
	"context"
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

func (s Store) InsertUser(ctx context.Context, u mongostore.User) (bool, error) {
	args := s.Called(ctx, u)
	return args.Get(0).(bool), args.Error(1)
}


// Close mocks close method
func (s Store) Close() error {
	args := s.Called()
	return args.Error(0)
}
