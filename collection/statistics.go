package collection

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
)

type (
	statistics struct {
		base
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
	_statistics := &statistics{}
	_statistics.exchange = exchange
	_statistics.queue = queue
	_statistics.base.subscribe = _statistics.subscribe
	return _statistics
}

func (c *statistics) validateRole(auth authorization) (string, error) {
	collection := facade.Db["collection_service"].Collection("authorization")
	var someone map[string]interface{}
	err = collection.FindOne(context.Background(), bson.D{
		{"appid", auth.Appid},
		{"secret", auth.Secret},
	}).Decode(&someone)

	return "", err
}

func (c *statistics) subscribe() {
	defer facade.WG.Done()

	for msg := range c.delivery {
		var source information
		if err = bson.UnmarshalExtJSON(msg.Body, true, &source); err != nil {
			c.ack(&msg)
			println(err.Error())
			continue
		}

		println(source.Authorization.Appid)

		//var app string
		//if _, err = c.validateRole(source.Authorization); err != nil {
		//	c.ack(&msg)
		//	println(err.Error())
		//	continue
		//}
		//
		//println(app)
	}
}
