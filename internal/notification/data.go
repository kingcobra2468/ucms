package notification

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

// Handle for Google FCM. Managing the sending of messages
// to the specified topic.
type AppAlerts interface {
	Connect(ctx context.Context) error
	SendNotification(title, body, url, serviceName string) error
}

// Client for message transmission onto FCM.
type Notifier struct {
	client *messaging.Client
	Topic  string
}

// Connect to Google's FCM.
func (n *Notifier) Connect(ctx context.Context) error {
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		return err
	}

	n.client, err = app.Messaging(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Send a notification onto the specified topic.
func (n *Notifier) SendNotification(title, body, url string) error {
	message := messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: url,
		},
		Topic: n.Topic,
	}
	_, err := n.client.Send(context.Background(), &message)

	return err
}
