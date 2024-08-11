package service

import (
	"binar/internal/app/queue/domain"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type queueService struct {
	conn   *amqp091.Connection
	logger *zap.Logger
}

func NewQueueService(conn *amqp091.Connection, logger *zap.Logger) domain.QueueService {
	return &queueService{
		conn:   conn,
		logger: logger,
	}
}

func (s *queueService) PublishMessage(message domain.QueueMessage) error {
	ch, err := s.conn.Channel()
	if err != nil {
		s.logger.Error("Failed to open a channel", zap.Error(err))
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		message.Type,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.logger.Error("Failed to declare a queue", zap.Error(err))
		return err
	}

	body, err := json.Marshal(message)
	if err != nil {
		s.logger.Error("Failed to marshal message", zap.Error(err))
		return err
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		s.logger.Error("Failed to publish a message", zap.Error(err))
		return err
	}

	return nil
}

func (s *queueService) ConsumeMessages(queueName string, handler func(domain.QueueMessage) error) error {
	ch, err := s.conn.Channel()
	if err != nil {
		s.logger.Error("Failed to open a channel", zap.Error(err))
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.logger.Error("Failed to declare a queue", zap.Error(err))
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		s.logger.Error("Failed to register a consumer", zap.Error(err))
		return err
	}

	for d := range msgs {
		var message domain.QueueMessage
		err := json.Unmarshal(d.Body, &message)
		if err != nil {
			s.logger.Error("Failed to unmarshal message", zap.Error(err))
			continue
		}

		err = handler(message)
		if err != nil {
			s.logger.Error("Failed to handle message", zap.Error(err))
		}
	}

	return nil
}
