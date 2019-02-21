package controller

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/uniplaces/carbon"
)

type (
	system struct {
		base
		systemDb string
	}

	logs struct {
		Publish string
		Data    map[string]interface{}
		Time    int64
	}
)

func NewSystem(database string, exchange string, queue string) *system {
	_system := &system{}
	_system.systemDb = database
	_system.exchange = exchange
	_system.queue = queue
	_system.base.subscribe = _system.subscribe
	return _system
}

func (c *system) validateWhitelist(value string) bool {
	collection := facade.Db[c.systemDb].Collection("whitelist")
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

		var _carbon *carbon.Carbon
		if _carbon, err = carbon.CreateFromTimestampUTC(source.Time); err != nil {
			println(err.Error())
			continue
		}

		source.Data["create_time"] = _carbon.Time
		collection := facade.Db[c.systemDb].Collection(source.Publish)

		if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
			println(err.Error())
		} else {
			c.ack(&msg)
		}
	}
}
