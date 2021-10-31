package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/kingcobra2468/ucms/internal/message"
	"github.com/kingcobra2468/ucms/internal/notification"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	n := notification.Notifier{Topic: "un"}
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
		url := fmt.Sprintf("%s:%d", "0.0.0.0", 8888)

		logger.Log("transport", "HTTP", "addr", url)
		errs <- http.ListenAndServe(url, h)
	}()

	logger.Log("exit", <-errs)
}
