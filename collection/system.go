package collection

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
	"time"
)

type (
	system struct {
		base
		database string
	}

	logs struct {
		Publish string
		Data    map[string]interface{}
		Time    int64
	}
)

func NewSystem(database string, exchange string, queue string) *system {
	_system := &system{}
	_system.database = database
	_system.exchange = exchange
	_system.queue = queue
	_system.base.subscribe = _system.subscribe
	return _system
}

func (c *system) validateWhitelist(value string) bool {
	collection := facade.Db[c.database].Collection("whitelist")
	var someone map[string]interface{}
	result := collection.FindOne(context.Background(), bson.D{{"domain", value}})
	return result.Decode(&someone) == nil
}

func (c *system) subscribe() {
	var err error
	defer facade.WG.Done()
	for msg := range c.delivery {
		var source logs
		if err = bson.UnmarshalExtJSON(msg.Body, true, &source); err != nil {
			c.ack(&msg)
			println(err.Error())
			continue
		}

		if !c.validateWhitelist(source.Publish) {
			continue
		}

		date := time.Unix(source.Time, 0)
		source.Data["create_time"] = date
		collection := facade.Db[c.database].Collection(source.Publish)

		if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
			println(err.Error())
		} else {
			c.ack(&msg)
		}
	}
}
