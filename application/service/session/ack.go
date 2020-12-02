package session

func (c *Session) Ack(queue string, receipt string) (err error) {
	if c.receipt.Empty(receipt) {
		return ReceiptHasExpired
	}
	receiptOption := c.receipt.Get(receipt)
	defer receiptOption.Channel.Close()
	if receiptOption.Queue != queue {
		return ReceiptIncorrect
	}
	if err = receiptOption.Delivery.Ack(false); err != nil {
		go c.logging(Log{
			Queue:   queue,
			Receipt: receipt,
			Payload: nil,
			Action:  "Ack",
		}, err)
		return
	}
	go c.logging(Log{
		Queue:   queue,
		Receipt: receipt,
		Payload: nil,
		Action:  "Ack",
	})
	return
}
