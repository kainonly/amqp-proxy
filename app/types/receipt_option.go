package types

import "github.com/streadway/amqp"

type ReceiptOption struct {
	Queue    string
	Delivery *amqp.Delivery
}
