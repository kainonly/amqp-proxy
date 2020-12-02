package session

import (
	"github.com/streadway/amqp"
)

type PublishOption struct {
	Exchange    string
	Key         string
	Mandatory   bool
	Immediate   bool
	ContentType string
	Body        []byte
}

func (c *Session) Publish(option PublishOption) (err error) {
	var channel *amqp.Channel
	if channel, err = c.conn.Channel(); err != nil {
		return
	}
	defer channel.Close()
	if err = channel.Publish(
		option.Exchange,
		option.Key,
		option.Mandatory,
		option.Immediate,
		amqp.Publishing{
			ContentType: option.ContentType,
			Body:        option.Body,
		},
	); err != nil {
		go c.logging(option, err)
		return
	}
	go c.logging(option)
	return
}
