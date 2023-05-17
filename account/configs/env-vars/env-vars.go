package envvars

import (
	"fmt"
	"time"

	"github.com/codingconcepts/env"
)

// EnvVars represents environment variables
type EnvVars struct {
	Service    Service
	HTTPServer HTTPServer
	Mongo      Mongo
}

// Service represents service configurations
type Service struct {
	Name          string `env:"SERVICE_NAME"`
	Environment   string `env:"SERVICE_ENV" default:"dev"`
	Release       string `env:"SERVICE_RELEASE"`
	ServiceBindIp string `env:"SERVICE_BIND_IP" default:"0.0.0.0"`
}

// Mongo represents mongo configurations
type Mongo struct {
	URI               string        `env:"MONGO_URI" required:"true"`
	Database          string        `env:"MONGO_DATABASE" required:"true"`
	ConnectTimeout    time.Duration `env:"MONGO_CONNECT_TIMEOUT" default:"10s"`
	PingTimeout       time.Duration `env:"MONGO_PING_TIMEOUT" default:"10s"`
	ReadTimeout       time.Duration `env:"MONGO_READ_TIMEOUT" default:"10s"`
	WriteTimeout      time.Duration `env:"MONGO_WRITE_TIMEOUT" default:"5s"`
	DisconnectTimeout time.Duration `env:"MONGO_DISCONNECT_TIMEOUT" default:"5s"`
}

// HTTPServer represents http server configurations
type HTTPServer struct {
	RestAddress     string        `env:"REST_HTTP_SERVER_ADDRESS" default:":8081"`
	ShutdownTimeout time.Duration `env:"HTTP_SERVER_SHUTDOWN_TIMEOUT" default:"10s"`
}

// LoadEnvVars loads and returns environment variables
func LoadEnvVars() (*EnvVars, error) {
	s := Service{}
	if err := env.Set(&s); err != nil {
		return nil, fmt.Errorf("loading service environment variables failed, %s", err.Error())
	}

	m := Mongo{}
	if err := env.Set(&m); err != nil {
		return nil, fmt.Errorf("loading mongo environment variables failed, %s", err.Error())
	}

	hs := HTTPServer{}
	if err := env.Set(&hs); err != nil {
		return nil, fmt.Errorf("loading http server environment variables failed, %s", err.Error())
	}

	ev := &EnvVars{
		Service:    s,
		HTTPServer: hs,
		Mongo:      m,
	}

	return ev, nil
}
