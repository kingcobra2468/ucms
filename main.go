package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"github.com/go-kit/kit/log"
	"github.com/kingcobra2468/ucms/internal/message"
	"github.com/kingcobra2468/ucms/internal/notification"
)

type config struct {
	ServiceHostname string `env:"UCMS_HOSTNAME" envDefault:"127.0.0.1"`
	ServicePort     int    `env:"UCMS_PORT" envDefault:"8080"`
	FcmTopic        string `env:"UCMS_FCM_TOPIC" envDefault:"un"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("%+v\n", err))
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	n := notification.Notifier{Topic: cfg.FcmTopic}
	{
		if err := n.Connect(context.Background()); err != nil {
			panic(err)
		}
	}

	var service message.MessageService = message.MessageBroadcast{Notifier: &n}
	service = message.LoggingMiddleware{Logger: logger, Next: service}
	var h http.Handler = message.MakeHTTPHandler(service)

	errs := make(chan error)
	// Listener for Ctrl+C signals.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	// Launch microservice.
	go func() {
		url := fmt.Sprintf("%s:%d", cfg.ServiceHostname, cfg.ServicePort)

		logger.Log("transport", "HTTP", "addr", url)
		errs <- http.ListenAndServe(url, h)
	}()

	logger.Log("exit", <-errs)
}
