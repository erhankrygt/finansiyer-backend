package main

import (
	"context"
	"errors"
	"finansiyer"
	envvars "finansiyer/configs/env-vars"
	"finansiyer/internal/service"
	mongostore "finansiyer/internal/store/mongo"
	httptransport "finansiyer/internal/transport/http"
	"github.com/go-kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "time", log.DefaultTimestampUTC)
	}

	var env *envvars.EnvVars
	var err error
	{
		env, err = envvars.LoadEnvVars()
		if err != nil {
			_ = logger.Log("error", err.Error())
			return
		}
	}

	var mongo mongostore.Store
	{
		mongo, err = mongostore.NewStore(env.Mongo)
		if err != nil {
			_ = logger.Log("error", err.Error())
			return
		}
	}

	var s finansiyer.Service
	{
		s = service.NewService(logger, mongo, env.Service.Environment)
	}

	var handler http.Handler
	{
		handler = httptransport.MakeHTTPHandler(log.With(logger, "transport", "http"), s)
	}

	// Rest Http Server struct with Handler and Addr
	var httpServer *http.Server
	{
		httpServer = &http.Server{
			Addr:    env.Service.ServiceBindIp + env.HTTPServer.RestAddress,
			Handler: handler,
		}
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- errors.New((<-c).String())
	}()

	// http Handler Serve with routine
	go func() {
		_ = logger.Log("transport", "http", "address", env.HTTPServer.RestAddress)

		err = httpServer.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			errs <- err
		}
	}()

	err = <-errs
	_ = logger.Log("error", err.Error())

	ctx, cf := context.WithTimeout(context.Background(), env.HTTPServer.ShutdownTimeout)
	defer cf()
	if err := httpServer.Shutdown(ctx); err != nil {
		_ = logger.Log("error", err.Error())
	}

	if err := mongo.Close(); err != nil {
		_ = logger.Log("error", err.Error())
	}
}
