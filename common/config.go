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
		SystemDatabase     string `ini:"system_database"`
		SystemExchange     string `ini:"system_exchange"`
		SystemQueue        string `ini:"system_queue"`
		StatisticsExchange string `ini:"statistics_exchange"`
		StatisticsQueue    string `ini:"statistics_queue"`
	}
)

func (m *Cogs) ValidateArgs() bool {
	return reflect.DeepEqual(m.Rabbitmq, rabbitmq{}) ||
		reflect.DeepEqual(m.Mongodb, mongodb{}) ||
		reflect.DeepEqual(m.Collection, collection{})
}

func (m *Cogs) RegisteredAMQP() error {
	var err error
	url := "amqp://" +
		m.Rabbitmq.Username + ":" +
		m.Rabbitmq.Password + "@" +
		m.Rabbitmq.Hostname + ":" +
		m.Rabbitmq.Port +
		m.Rabbitmq.Vhost

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

func (m *Cogs) RegisteredMongo() error {
	var err error
	dsn := "mongodb://" +
		m.Mongodb.Username + ":" +
		m.Mongodb.Password + "@" +
		m.Mongodb.Hostname + ":" +
		m.Mongodb.Port

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
	facade.Db["collection_service"] = facade.MGOClient.Database("collection_service")
	facade.Db[m.Collection.SystemDatabase] = facade.MGOClient.Database(m.Collection.SystemDatabase)
	return nil
}
