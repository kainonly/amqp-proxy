package facade

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/streadway/amqp"
)

var MGOClient *mongo.Client

var MGODb map[string]*mongo.Database

var AMQPConnection *amqp.Connection

var AMQPChannel *amqp.Channel
