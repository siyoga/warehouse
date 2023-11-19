package broker

import (
	"fmt"
	"warehouseai/mail/config"

	rmq "github.com/rabbitmq/amqp091-go"
)

func NewMailConsumer() (*rmq.Channel, rmq.Queue, *rmq.Connection) {
	config := config.NewMailBrokerCfg()

	conn, err := rmq.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.User, config.Password, config.Host, config.Port))

	if err != nil {
		panic(fmt.Sprintf("Unable to open connect to RabbitMQ; %s", err))
	}

	ch, err := conn.Channel()

	if err != nil {
		panic(fmt.Sprintf("Unable to open channel; %s", err))
	}

	queue, err := ch.QueueDeclare(
		"mail",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(fmt.Sprintf("Unable to create queue in the channel; %s", err))
	}

	return ch, queue, conn
}
