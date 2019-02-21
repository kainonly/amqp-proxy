package collection

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
	"time"
)

type (
	system struct {
		database string
		common
	}

	logs struct {
		Publish string
		Data    map[string]interface{}
		Time    int64
	}
)

func NewSystem(database string, exchange string, queue string) *system {
	m := &system{}
	m.database = database
	m.exchange = exchange
	m.queue = queue
	return m
}

func (m *system) validateWhitelist(value string) bool {
	collection := facade.Db[m.database].Collection("whitelist")
	var someone map[string]interface{}
	result := collection.FindOne(context.Background(), bson.D{{"domain", value}})
	return result.Decode(&someone) == nil
}

func (m *system) Run() {
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

func (m *system) subscribe() {
	var err error
	defer facade.WG.Done()
	for msg := range m.delivery {
		var source logs
		if err = bson.UnmarshalExtJSON(msg.Body, true, &source); err != nil {
			m.ack(&msg)
			println(err.Error())
			continue
		}

		if !m.validateWhitelist(source.Publish) {
			continue
		}

		date := time.Unix(source.Time, 0)
		source.Data["create_time"] = date
		collection := facade.Db[m.database].Collection(source.Publish)

		if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
			println(err.Error())
		} else {
			m.ack(&msg)
		}
	}
}
