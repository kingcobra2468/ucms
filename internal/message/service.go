package message

type MessageService interface {
	push(message string) (bool, error)
}

type Message struct {
}

func (m Message) push(title, body, url, serviceName string) (bool, error) {
	return true, nil
}
