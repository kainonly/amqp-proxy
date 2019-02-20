package collection

import (
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/bson"
)

type (
	statistics struct {
		common
	}

	information struct {
		authorization
		Data      map[string]interface{}
		TimeField []string `bson:"time_field" json:"time_field"`
	}

	authorization struct {
		Appid  string
		Secret string
	}
)

func NewStatistics(exchange string, queue string) *statistics {
	m := &statistics{}
	m.exchange = exchange
	m.queue = queue
	return m
}

func (m *statistics) _ValidateRole(auth authorization) (string, error) {
	return "", nil
}

func (m *statistics) Run() {
	if err = m.defined(); err != nil {
		panic(err.Error())
	}

	if err = facade.AMQPChannel.Qos(1, 0, false); err != nil {
		panic(err.Error())
	}

	// start consume
	if m.delivery, err = facade.AMQPChannel.Consume(
		m.queue,
		"",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic(err.Error())
	}

	go m.subscribe()
}

func (m *statistics) subscribe() {
	var err error
	defer facade.WG.Done()

	for msg := range m.delivery {
		var source information
		if err = bson.UnmarshalExtJSON(msg.Body, true, &source); err != nil {
			m.ack(&msg)
			println(err.Error())
			continue
		}

		var database string
		if database, err = m._ValidateRole(source.authorization); err != nil {
			m.ack(&msg)
			println(err.Error())
			continue
		}

		println(database)
		m.ack(&msg)
	}
}
