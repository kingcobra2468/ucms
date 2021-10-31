package message

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Schema for new notification request
type NewNotificationRequest struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	ImageURL    string `json:"image,omitempty"`
	ServiceName string `json:"service_name"`
}

// Schema for new notification response
type NewNotificationResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error,omitempty"`
}

// Endpoint for sending new notification
func makeNewNotificationEndpoint(ms MessageService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NewNotificationRequest)
		if err := ms.Push(req.Title, req.Body, req.ImageURL, req.ServiceName); err != nil {
			return NewNotificationResponse{Success: false, Error: err}, nil
		}

		return NewNotificationResponse{Success: true, Error: nil}, nil
	}
}
