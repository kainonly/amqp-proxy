package common

import (
	"context"
	"github.com/kainonly/collection-service/src/facade"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/streadway/amqp"
	"reflect"
	"time"
)

type Cogs struct {
	Rabbitmq `ini:"rabbitmq"`
	Mongodb  `ini:"mongodb"`
	Logs     `ini:"logs"`
}

type Rabbitmq struct {
	Hostname string `ini:"hostname"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Vhost    string `ini:"vhost"`
}

type Mongodb struct {
	Hostname string `ini:"hostname"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type Logs struct {
	SystemDatabase string `ini:"system_database"`
	SystemExchange string `ini:"system_exchange"`
	SystemQueue    string `ini:"system_queue"`
}

func (m *Cogs) CheckArgs() bool {
	return reflect.DeepEqual(m.Rabbitmq, Rabbitmq{}) ||
		reflect.DeepEqual(m.Mongodb, Mongodb{}) ||
		reflect.DeepEqual(m.Logs, Logs{})
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
	defer facade.AMQPConnection.Close()

	// Create AMQP channel
	if facade.AMQPChannel, err = facade.AMQPConnection.Channel(); err != nil {
		return err
	}
	defer facade.AMQPChannel.Close()

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
	defer cancel()
	if err = facade.MGOClient.Connect(ctx); err != nil {
		return err
	}

	// using database
	facade.MGODb[m.SystemDatabase] = facade.MGOClient.Database(m.SystemDatabase)
	return nil
}
