package session

import "errors"

func (c *Session) Ack(queue string, receipt string) (err error) {
	msg := c.receipt.Get(receipt)
	if msg == nil {
		return errors.New("the receipt has expired")
	}
	if msg.Queue != queue {
		return errors.New("the receipt verification is incorrect")
	}
	err = msg.Delivery.Ack(false)
	if err != nil {
		return
	}
	return
}
