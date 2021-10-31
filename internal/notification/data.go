package notification

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type AppAlerts interface {
	Connect(ctx context.Context) error
	SendNotification(title, body, url, serviceName string) error
}

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

func (n *Notifier) SendNotification(title, body, url string) error {
	message := messaging.Message{
		Notification: &messaging.Notification{
			Title:    title,
			Body:     body,
			ImageURL: url,
		},
		Topic: n.Topic,
		Webpush: &messaging.WebpushConfig{
			Notification: &messaging.WebpushNotification{
				Icon: url,
			},
		},
	}
	_, err := n.client.Send(context.Background(), &message)

	return err
}
