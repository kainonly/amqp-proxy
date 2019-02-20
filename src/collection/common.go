package collection

import (
	"github.com/kainonly/collection-service/src/facade"
	"github.com/streadway/amqp"
)

type common struct {
	exchange string
	queue    string
	delivery <-chan amqp.Delivery
}

var err error

func (m *common) defined() error {
	// declare exchange
	if err = facade.AMQPChannel.ExchangeDeclare(
		m.exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// declare queue
	if _, err = facade.AMQPChannel.QueueDeclare(
		m.queue,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// bind queue
	if err = facade.AMQPChannel.QueueBind(
		m.queue,
		"",
		m.exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (m *common) ack(msg *amqp.Delivery) {
	if err = msg.Ack(false); err != nil {
		panic(err.Error())
	}
}
