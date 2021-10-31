package message

import (
	"fmt"

	"github.com/kingcobra2468/ucms/internal/notification"
)

// Service for endpoints related to notification and message
// transmission onto FCM.
type MessageService interface {
	Push(title, body, url, serviceName string) error
}

// Driver for notification and message transmission.
type MessageBroadcast struct {
	Notifier *notification.Notifier
}

// Push a notification onto a FCM topic.
func (mb MessageBroadcast) Push(title, body, url, serviceName string) error {
	err := mb.Notifier.SendNotification(fmt.Sprintf("%s - %s", serviceName, title), body, url)
	return err
}
