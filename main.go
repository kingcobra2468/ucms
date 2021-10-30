package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/kingcobra2468/ucms/internal/message"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//var ds notification.DeviceSubscriber = notification.DeviceSubscriber{Topic: *topic}
	//{
	//	ds.Connect(context.Background())
	//}

	var service message.Message = message.Message{}
	//service = device.LoggingMiddleware{Logger: logger, Next: service}
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
