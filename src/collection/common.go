package collection

import "github.com/kainonly/collection-service/src/facade"

type Common struct {
	Exchange string
	Queue    string
}

func (m *Common) _DeclareMQ() error {
	// declare exchange
	if err := facade.AMQPChannel.ExchangeDeclare(
		m.Exchange,
		"direct",
		false,
		true,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// declare queue
	if _, err := facade.AMQPChannel.QueueDeclare(
		m.Queue,
		false,
		true,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	// bind queue
	if err := facade.AMQPChannel.QueueBind(
		m.Queue,
		"",
		m.Exchange,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}