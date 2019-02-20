package collection

import (
	"context"
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/streadway/amqp"
	"time"
)

type System struct {
	Database string
	Common
}

type Logs struct {
	Publish string
	Data    map[string]interface{}
	Time    int64
}

func NewSystem(database string, exchange string, queue string) *System {
	system := &System{}
	system.Database = database
	system.Exchange = exchange
	system.Queue = queue
	return system
}

func (m *System) _ValidateWhitelist(value string) bool {
	collection := facade.Db[m.Database].Collection("whitelist")
	var someone map[string]interface{}
	result := collection.FindOne(context.Background(), bson.D{{"domain", value}})
	return result.Decode(&someone) == nil
}

func (m *System) Subscribe() {
	var err error
	if err = m._DeclareMQ(); err != nil {
		panic(err.Error())
	}

	if err = facade.AMQPChannel.Qos(1, 0, false); err != nil {
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
			var source Logs
			if err = bson.UnmarshalExtJSON(x.Body, true, &source); err != nil {
				panic(err.Error())
			}

			if m._ValidateWhitelist(source.Publish) {
				date := time.Unix(source.Time, 0)
				source.Data["create_time"] = date
				collection := facade.Db[m.Database].Collection(source.Publish)
				if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
					panic(err.Error())
				}
			}

			if err = x.Ack(false); err != nil {
				panic(err.Error())
			}
		}
	}()
}
