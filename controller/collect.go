package controller

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/uniplaces/carbon"
	"reflect"
)

type (
	collect struct {
		base
		app string
	}

	information struct {
		Authorization authorization
		Motivation    string
		Data          map[string]interface{}
		Time_Field    []string
	}

	authorization struct {
		Appid  string
		Secret string
	}
)

func NewStatistics(database string, exchange string, queue string) *collect {
	_collect := &collect{}
	_collect.database = database
	_collect.exchange = exchange
	_collect.queue = queue
	_collect.base.subscribe = _collect.subscribe
	return _collect
}

func (c *collect) validateRole(auth authorization) bool {
	collection := facade.Db[c.database].Collection("authorization")
	var someone map[string]interface{}
	if err = collection.FindOne(context.Background(), bson.D{
		{"appid", auth.Appid},
		{"secret", auth.Secret},
	}).Decode(&someone); err == nil {
		c.app = someone["app"].(string)
		return true
	} else {
		return false
	}
}

func (c *collect) subscribe() {
	defer facade.WG.Done()

	for msg := range c.delivery {
		var source information
		if err = bson.UnmarshalExtJSON(msg.Body, true, &source); err != nil {
			c.ack(&msg)
			println(err.Error())
			continue
		}

		if !c.validateRole(source.Authorization) {
			c.ack(&msg)
			println("authorization failed!")
			continue
		}

		if facade.Db[c.app] == nil {
			facade.Db[c.app] = facade.MGOClient.Database(c.app)
		}

		for _, x := range source.Time_Field {
			if _, ok := source.Data[x]; ok {
				if reflect.TypeOf(source.Data[x]).String() != "int32" {
					continue
				}

				var _carbon *carbon.Carbon
				if _carbon, err = carbon.CreateFromTimestampUTC(int64(source.Data[x].(int32))); err != nil {
					println(err.Error())
					source.Data[x] = nil
				} else {
					source.Data[x] = _carbon.Time
				}
			}
		}

		collection := facade.Db[c.app].Collection(source.Motivation)
		if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
			println(err.Error())
		} else {
			c.ack(&msg)
		}
	}
}
