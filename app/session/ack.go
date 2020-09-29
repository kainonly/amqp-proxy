package session

import (
	"errors"
)

func (c *Session) Ack(queue string, receipt string) (err error) {
	receiptOption := c.receipt.Get(receipt)
	if receiptOption == nil {
		err = errors.New("the receipt has expired")
		c.collectFromAction(queue, receipt, nil, "Ack", err)
		return
	}
	if receiptOption.Queue != queue {
		err = errors.New("the receipt verification is incorrect")
		c.collectFromAction(queue, receipt, nil, "Ack", err)
		return
	}
	err = receiptOption.Delivery.Ack(false)
	if err != nil {
		c.collectFromAction(queue, receipt, nil, "Ack", err)
		return
	}
	receiptOption.Channel.Close()
	c.collectFromAction(queue, receipt, nil, "Ack", nil)
	return
}
