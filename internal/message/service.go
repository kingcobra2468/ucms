package message

import (
	"fmt"

	"github.com/kingcobra2468/ucms/internal/notification"
)

type MessageService interface {
	Push(title, body, url, serviceName string) (bool, error)
}

type MessageBroadcast struct {
	Notifier *notification.Notifier
}

func (mb MessageBroadcast) Push(title, body, url, serviceName string) (bool, error) {
	mb.Notifier.SendNotification(fmt.Sprintf("%s - %s", serviceName, title), body, url)
	return true, nil
}
