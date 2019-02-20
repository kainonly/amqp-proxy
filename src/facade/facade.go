package facade

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/streadway/amqp"
)

var MGOClient *mongo.Client

var Db map[string]*mongo.Database

var Cancel context.CancelFunc

var AMQPConnection *amqp.Connection

var AMQPChannel *amqp.Channel
