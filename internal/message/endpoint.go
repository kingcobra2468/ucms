package message

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type NewNotificationRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	ImageURL    string `json:"image,omitempty"`
	ServiceName string `json:"service_name"`
}

type NewNotificationResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error,omitempty"`
}

func makeNewNotificationEndpoint(ms MessageService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewNotificationRequest)
		ms.Push(req.Title, req.Body, req.ImageURL, req.ServiceName)

		return NewNotificationResponse{Success: true, Error: nil}, nil
	}
}
