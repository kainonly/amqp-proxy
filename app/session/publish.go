package session

import (
	"amqp-proxy/app/types"
	jsoniter "github.com/json-iterator/go"
	"github.com/streadway/amqp"
	"github.com/xeipuuv/gojsonschema"
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
	var message map[string]interface{}
	if err != nil {
		message = map[string]interface{}{
			"Topic": option.Exchange,
			"Key":   option.Key,
			"Message": map[string]string{
				"errs": err.Error(),
			},
			"Status": false,
			"Time":   time.Now().Unix(),
		}
	} else {
		var body interface{}
		result, err := gojsonschema.Validate(
			gojsonschema.NewBytesLoader([]byte(`{"type":"object"}`)),
			gojsonschema.NewBytesLoader(option.Body),
		)
		if err != nil {
			body = map[string]interface{}{
				"raw": string(option.Body),
			}
		} else {
			if result.Valid() {
				jsoniter.Unmarshal(option.Body, &body)
			} else {
				body = map[string]interface{}{
					"raw": string(option.Body),
				}
			}
		}
		message = map[string]interface{}{
			"Topic":   option.Exchange,
			"Key":     option.Key,
			"Message": body,
			"Status":  true,
			"Time":    time.Now().Unix(),
		}
	}
	c.logging.Push(c.pipe.PublishID, message)
	return
}
