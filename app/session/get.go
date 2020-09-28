package session

import (
	"amqp-proxy/app/types"
	"errors"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

func (c *Session) Get(queue string) (receipt string, body []byte, err error) {
	var channel *amqp.Channel
	channel, err = c.conn.Channel()
	if err != nil {
		go c.collectFromAction(queue, nil, nil, "Get", err)
		return
	}
	notifyClose := make(chan *amqp.Error)
	channel.NotifyClose(notifyClose)
	msg, ok, err := channel.Get(queue, false)
	if err != nil {
		go c.collectFromAction(queue, nil, nil, "Get", err)
		return
	}
	if ok == false {
		err = errors.New("available queue does not exist")
		channel.Close()
		go c.collectFromAction(queue, nil, nil, "Get", err)
		return
	}
	receipt = uuid.New().String()
	c.receipt.Set(receipt, &types.ReceiptOption{
		Queue:    queue,
		Channel:  channel,
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
	go c.collectFromAction(queue, receipt, string(body), "Get", nil)
	return
}
