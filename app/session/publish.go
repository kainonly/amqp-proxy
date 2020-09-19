package session

import (
	"amqp-proxy/app/types"
	"github.com/streadway/amqp"
	"time"
)

func (c *Session) Publish(option *types.PublishOption) (err error) {
	var channel *amqp.Channel
	channel, err = c.conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()
	err = channel.Publish(
		option.Exchange,
		option.Key,
		option.Mandatory,
		option.Immediate,
		amqp.Publishing{
			ContentType: option.ContentType,
			Body:        option.Body,
		},
	)
	if err != nil {
		return
	}
	c.logging.Push(c.pipe.Publish, map[string]interface{}{
		"Topic":   option.Exchange,
		"Key":     option.Key,
		"Payload": string(option.Body),
		"Time":    time.Now().Unix(),
	})
	return
}
