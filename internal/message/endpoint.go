package message

import "firebase.google.com/go/messaging"

type NewNotificationRequest struct {
	messaging.Notification
	ServiceName string `json:"service_name"`
}

type NewNotificationResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error,omitempty"`
}

func makeNewNotificationEndpoint() {}
