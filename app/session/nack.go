package session

import "errors"

func (c *Session) Nack(queue string, receipt string) (err error) {
	receiptOption := c.receipt.Get(receipt)
	if receiptOption == nil {
		return errors.New("the receipt has expired")
	}
	if receiptOption.Queue != queue {
		return errors.New("the receipt verification is incorrect")
	}
	err = receiptOption.Delivery.Nack(false, false)
	if err != nil {
		return
	}
	receiptOption.Channel.Close()
	return
}
