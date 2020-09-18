package session

import (
	"amqp-proxy/app/types"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func (c *Session) Get(queue string) (data types.Data, err error) {
	var channel *amqp.Channel
	channel, err = c.conn.Channel()
	if err != nil {
		return
	}
	msg, _, err := channel.Get(queue, false)
	if err != nil {
		return
	}
	receipt := uuid.New().String()
	c.delivery.Set(receipt, &msg)
	data = types.Data{
		Receipt: receipt,
		Body:    msg.Body,
	}
	return
}
