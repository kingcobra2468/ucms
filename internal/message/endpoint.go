package message

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type NewNotificationRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	ImageURL    string `json:"image"`
	ServiceName string `json:"service_name"`
}

type NewNotificationResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error,omitempty"`
}

func makeNewNotificationEndpoint(ms MessageService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewNotificationRequest)
		fmt.Println(req)

		return NewNotificationResponse{Success: true, Error: nil}, nil
	}
}
