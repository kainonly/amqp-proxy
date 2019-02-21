package common

import (
	"context"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/streadway/amqp"
	"reflect"
	"time"
)

type (
	Cogs struct {
		Rabbitmq   rabbitmq   `ini:"rabbitmq"`
		Mongodb    mongodb    `ini:"mongodb"`
		Collection collection `ini:"collection"`
	}

	rabbitmq struct {
		Hostname string `ini:"hostname"`
		Port     string `ini:"port"`
		Username string `ini:"username"`
		Password string `ini:"password"`
		Vhost    string `ini:"vhost"`
	}

	mongodb struct {
		Hostname string `ini:"hostname"`
		Port     string `ini:"port"`
		Username string `ini:"username"`
		Password string `ini:"password"`
	}

	collection struct {
		Database       string `ini:"database"`
		Exchange       string `ini:"exchange"`
		Queue          string `ini:"queue"`
		SystemExchange string `ini:"system_exchange"`
		SystemQueue    string `ini:"system_queue"`
	}
)

func (c *Cogs) ValidateArgs() bool {
	return reflect.DeepEqual(c.Rabbitmq, rabbitmq{}) ||
		reflect.DeepEqual(c.Mongodb, mongodb{}) ||
		reflect.DeepEqual(c.Collection, collection{})
}

func (c *Cogs) RegisteredAMQP() error {
	var err error
	url := "amqp://" +
		c.Rabbitmq.Username + ":" +
		c.Rabbitmq.Password + "@" +
		c.Rabbitmq.Hostname + ":" +
		c.Rabbitmq.Port +
		c.Rabbitmq.Vhost

	// Connect RabbitMQ
	if facade.AMQPConnection, err = amqp.Dial(url); err != nil {
		return err
	}

	// Create AMQP channel
	if facade.AMQPChannel, err = facade.AMQPConnection.Channel(); err != nil {
		return err
	}

	return nil
}

func (c *Cogs) RegisteredMongo() error {
	var err error
	dsn := "mongodb://" +
		c.Mongodb.Username + ":" +
		c.Mongodb.Password + "@" +
		c.Mongodb.Hostname + ":" +
		c.Mongodb.Port

	// create mongodb client
	if facade.MGOClient, err = mongo.NewClient(dsn); err != nil {
		return err
	}

	// connect mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = facade.MGOClient.Connect(ctx); err != nil {
		return err
	}
	defer cancel()

	// using database
	facade.Db[c.Collection.Database] = facade.MGOClient.Database(c.Collection.Database)
	return nil
}
