package collection

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
)

type (
	statistics struct {
		common
	}

	information struct {
		Authorization authorization
		Data          map[string]interface{}
		Time_Field    []string
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

func (m *statistics) validateRole(auth authorization) (string, error) {
	collection := facade.Db["collection_service"].Collection("authorization")
	var someone map[string]interface{}
	err = collection.FindOne(context.Background(), bson.D{
		{"appid", auth.Appid},
		{"secret", auth.Secret},
	}).Decode(&someone)

	return "", err
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
	//var err error
	defer facade.WG.Done()

	for msg := range m.delivery {
		var source information
		if err = bson.UnmarshalExtJSON(msg.Body, true, &source); err != nil {
			//m.ack(&msg)
			println(err.Error())
			continue
		}

		println(source.Authorization.Appid)
		println(source.Time_Field)
		//
		////var app string
		//if _, err = m.validateRole(source.authorization); err != nil {
		//	m.ack(&msg)
		//	println(err.Error())
		//	continue
		//}
		//
		////println(app)
		m.ack(&msg)
	}
}
