package facade

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/streadway/amqp"
	"sync"
)

var (
	MGOClient      *mongo.Client
	Db             map[string]*mongo.Database
	AMQPConnection *amqp.Connection
	AMQPChannel    *amqp.Channel
	WG             sync.WaitGroup
)
