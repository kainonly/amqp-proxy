package controller

import (
	"github.com/kainonly/collection-service/facade"
	"github.com/streadway/amqp"
)

type (
	base struct {
		exchange  string
		queue     string
		delivery  <-chan amqp.Delivery
		subscribe func()
	}
)

var err error

func (c *base) defined() error {
	// declare exchange
	if err = facade.AMQPChannel.ExchangeDeclare(
		c.exchange,
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
		c.queue,
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
		c.queue,
		"",
		c.exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *base) Run() {
	if err = c.defined(); err != nil {
		panic(err.Error())
	}

	if err = facade.AMQPChannel.Qos(1, 0, false); err != nil {
		panic(err.Error())
	}

	// start consume
	if c.delivery, err = facade.AMQPChannel.Consume(
		c.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err.Error())
	}

	if c.subscribe != nil {
		go c.subscribe()
	}
}

func (c *base) ack(msg *amqp.Delivery) {
	if err = msg.Ack(false); err != nil {
		panic(err.Error())
	}
}
