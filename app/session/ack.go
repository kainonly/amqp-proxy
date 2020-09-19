package session

import (
	"errors"
	"time"
)

func (c *Session) Ack(queue string, receipt string) (err error) {
	receiptOption := c.receipt.Get(receipt)
	if receiptOption == nil {
		return errors.New("the receipt has expired")
	}
	if receiptOption.Queue != queue {
		return errors.New("the receipt verification is incorrect")
	}
	err = receiptOption.Delivery.Ack(false)
	if err != nil {
		return
	}
	receiptOption.Channel.Close()
	c.logging.Push(c.pipe.Message, map[string]interface{}{
		"Queue":   queue,
		"Receipt": receipt,
		"Payload": nil,
		"Action":  "Ack",
		"Time":    time.Now().Unix(),
	})
	return
}
