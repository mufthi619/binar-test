package domain

type QueueMessage struct {
	Type    string
	Payload []byte
}

type QueueService interface {
	PublishMessage(message QueueMessage) error
	ConsumeMessages(queueName string, handler func(QueueMessage) error) error
}
