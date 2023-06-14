package mongostore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"
	envvars "wallet/configs/env-vars"
)

// collections
const (
	userCollection = "users"
)

// Store defines behaviors of mongo store
type Store interface {
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

func (s store) Close() error {
	ctx, cf := context.WithTimeout(context.Background(), s.disconnectTimeout)
	defer cf()

	return s.c.Disconnect(ctx)
}
