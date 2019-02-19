package logs

import (
	"context"
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/streadway/amqp"
	"time"
)

type System struct {
	Database string
	Exchange string
	Queue    string
}

type Logs struct {
	Publish string
	Data    map[string]interface{}
	Time    int64
}

func NewSystem(database string, exchange string, queue string) *System {
	return &System{
		Database: database,
		Exchange: exchange,
		Queue:    queue,
	}
}

func (m *System) Subscribe() {
	// declare exchange
	if err = facade.AMQPChannel.ExchangeDeclare(
		m.Exchange,
		"direct",
		false,
		true,
		false,
		false,
		nil,
	); err != nil {
		panic(err.Error())
	}

	// declare queue
	if _, err = facade.AMQPChannel.QueueDeclare(
		m.Queue,
		false,
		true,
		false,
		false,
		nil,
	); err != nil {
		panic(err.Error())
	}

	if err = facade.AMQPChannel.QueueBind(
		m.Queue,
		"",
		m.Exchange,
		false,
		nil,
	); err != nil {
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

	// subscribe
	for x := range msg {
		var source Logs
		if err = bson.UnmarshalExtJSON(x.Body, true, &source); err != nil {
			panic(err.Error())
		}

		if CheckAllowDomain(source.Publish) {
			date := time.Unix(source.Time, 0)
			source.Data["create_time"] = date
			collection := facade.MGODb[m.Database].Collection(source.Publish)
			if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
				panic(err.Error())
			}
		}

		if err = facade.AMQPChannel.Ack(x.DeliveryTag, false); err != nil {
			panic(err.Error())
		}
	}
}
