package session

import (
	"amqp-proxy/app/types"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func (c *Session) Get(queue string) (receipt string, body []byte, err error) {
	var channel *amqp.Channel
	channel, err = c.conn.Channel()
	if err != nil {
		return
	}
	notifyClose := make(chan *amqp.Error)
	channel.NotifyClose(notifyClose)
	msg, _, err := channel.Get(queue, false)
	if err != nil {
		return
	}
	receipt = uuid.New().String()
	c.receipt.Set(receipt, &types.ReceiptOption{
		Queue:    queue,
		Delivery: &msg,
	})
	body = msg.Body
	go func() {
		select {
		case <-notifyClose:
			c.receipt.Delete(receipt)
			break
		}
	}()
	return
}
