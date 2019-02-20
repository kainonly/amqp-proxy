package collection

import (
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/bson"
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

func (m *Statistics) _ValidateRole(authorization Authorization) (string, error) {
	return "", nil
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
			var source Information
			if err = bson.UnmarshalExtJSON(x.Body, true, &source); err != nil {
				panic(err.Error())
			}

			var database string
			if database, err = m._ValidateRole(source.Authorization); err != nil {
				panic(err.Error())
			}

			println(database)
			println(x.Body)
		}
	}()
}
