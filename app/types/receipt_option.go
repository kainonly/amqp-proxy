package types

import "github.com/streadway/amqp"

type ReceiptOption struct {
	Queue    string
	Channel  *amqp.Channel
	Delivery *amqp.Delivery
}
