package collection

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/bson"
)

type (
	statistics struct {
		base
		app string
	}

	information struct {
		Authorization authorization
		Namespace     string
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

func (c *statistics) validateRole(auth authorization) bool {
	collection := facade.Db["collection_service"].Collection("authorization")
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

func (c *statistics) subscribe() {
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

		collection := facade.Db[c.app].Collection(source.Namespace)

		if _, err = collection.InsertOne(context.Background(), source.Data); err != nil {
			println(err.Error())
		} else {
			c.ack(&msg)
		}
	}
}
