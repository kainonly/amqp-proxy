package collection

import (
	"github.com/kainonly/collection-service/src/facade"
	"github.com/streadway/amqp"
)

type Statistics struct {
	Common
}

type Information struct {
	Authorization
	Data map[string]interface{}
}

type Authorization struct {
	Appid  string
	Secret string
}

func NewStatistics(exchange string, queue string) *Statistics {
	statistics := &Statistics{}
	statistics.Exchange = exchange
	statistics.Queue = queue
	return statistics
}

func (m *Statistics) _ValidateRole() {

}

func (m *Statistics) Subscribe() {
	var err error
	if err = m._DeclareMQ(); err != nil {
		panic(err.Error())
	}

	// start consume
	var msg <-chan amqp.Delivery
	if msg, err = facade.AMQPChannel.Consume(
		m.Queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err.Error())
	}

	go func() {
		defer facade.WG.Done()
		for x := range msg {
			println(x.Body)
		}
	}()
}
