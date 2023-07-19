package mongostore

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
	envvars "wallet/configs/env-vars"
)

// collections
const (
	userCollection        = "users"
	bankCollection        = "banks"
	bankAccountCollection = "bankAccounts"
)

// errors
var (
	ErrNoDocumentsFound       = errors.New("no documents found")
	ErrDecodingDocumentRecord = errors.New("decoding document record failed")
	ErrNotFoundDocument       = errors.New("not found document failed")
	ErrBankCouldNotCreate     = errors.New("bank could not create")
)

// Store defines behaviors of mongo store
type Store interface {
	GetBanks(ctx context.Context, u User) ([]Bank, error)
	CreateBank(ctx context.Context, b Bank) (bool, error)
	CreateBankAccount(ctx context.Context, ba BankAccount) (bool, error)
	Close() error
}

// store represents mongo store
type store struct {
	uri               string
	database          string
	connectTimeout    time.Duration
	pingTimeout       time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	disconnectTimeout time.Duration
	c                 *mongo.Client
	db                *mongo.Database
	mtx               sync.Mutex
}

// NewStore creates and returns mongo store
func NewStore(m envvars.Mongo) (Store, error) {
	s := &store{
		uri:               m.URI,
		database:          m.Database,
		connectTimeout:    m.ConnectTimeout,
		readTimeout:       m.ReadTimeout,
		writeTimeout:      m.WriteTimeout,
		pingTimeout:       m.PingTimeout,
		disconnectTimeout: m.DisconnectTimeout,
	}

	cctx, ccf := context.WithTimeout(context.Background(), s.connectTimeout)
	defer ccf()

	opts := options.Client()
	opts.ApplyURI(s.uri)

	c, err := mongo.Connect(cctx, opts)
	if err != nil {
		return nil, fmt.Errorf("connecting failed, %s", err.Error())
	}

	s.c = c

	pctx, pcf := context.WithTimeout(context.Background(), s.pingTimeout)
	defer pcf()

	if err := s.c.Ping(pctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("pinging failed, %s", err.Error())
	}

	s.db = c.Database(s.database)

	return s, nil
}

func (s store) GetBanks(ctx context.Context, u User) ([]Bank, error) {
	d := s.db.Collection(bankCollection)

	rctx, rcf := context.WithTimeout(ctx, s.readTimeout)
	defer rcf()

	res, err := d.Find(rctx, bson.M{})
	if err != nil {
		return nil, ErrDecodingDocumentRecord
	}

	dl := make([]Bank, 0)

	for res.Next(rctx) {
		doc := Bank{}
		err = res.Decode(&doc)
		if err != nil {
			return nil, ErrNotFoundDocument
		}

		dl = append(dl, doc)
	}

	return dl, nil
}

func (s store) CreateBank(ctx context.Context, b Bank) (bool, error) {

	c := s.db.Collection(bankCollection)

	wctx, wcf := context.WithTimeout(ctx, s.writeTimeout)
	defer wcf()

	upc := Bank{
		Title:     b.Title,
		WebSite:   b.WebSite,
		IsActive:  true,
		IsDeleted: false,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	_, err := c.InsertOne(wctx, upc)
	if err != nil {
		return false, ErrBankCouldNotCreate
	}

	return true, nil
}

func (s store) CreateBankAccount(ctx context.Context, ba BankAccount) (bool, error) {

	c := s.db.Collection(bankAccountCollection)

	wctx, wcf := context.WithTimeout(ctx, s.writeTimeout)
	defer wcf()

	upc := BankAccount{
		Bank: Bank{
			User: User{},
		},
		IBAN:          ba.IBAN,
		AccountNumber: ba.AccountNumber,
		IsActive:      true,
		IsDeleted:     false,
		CreatedAt:     primitive.NewDateTimeFromTime(time.Now().UTC()),
	}

	_, err := c.InsertOne(wctx, upc)
	if err != nil {
		return false, ErrBankCouldNotCreate
	}

	return true, nil
}

func (s store) Close() error {
	ctx, cf := context.WithTimeout(context.Background(), s.disconnectTimeout)
	defer cf()

	return s.c.Disconnect(ctx)
}
