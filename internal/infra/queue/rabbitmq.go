package queue

import (
	"binar/pkg/config"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewRabbitMQConn(config *config.Config) (conn *amqp.Connection, err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		config.RabbitConfig.User,
		config.RabbitConfig.Password,
		config.RabbitConfig.Host,
		config.RabbitConfig.Port,
	)
	conn, err = amqp.Dial(url)
	return
}
