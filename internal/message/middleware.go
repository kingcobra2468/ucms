package message

import (
	"time"

	"github.com/go-kit/kit/log"
)

// Middlewere for performing request-based logging of the endpoints.
type LoggingMiddleware struct {
	Logger log.Logger
	Next   MessageService
}

// Logging wrapper for token registration logic.
func (lm LoggingMiddleware) Push(title, body, url, serviceName string) (bool, error) {
	defer func(begin time.Time) {
		lm.Logger.Log(
			"method", "Push",
			"took", time.Since(begin),
		)
	}(time.Now())

	status, err := lm.Next.Push(title, body, url, serviceName)
	return status, err
}
