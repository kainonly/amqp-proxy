package facade

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/streadway/amqp"
	"reflect"
	"runtime"
	"sync"
)

var (
	MGOClient      *mongo.Client
	Db             map[string]*mongo.Database
	AMQPConnection *amqp.Connection
	AMQPChannel    *amqp.Channel
	WG             sync.WaitGroup
)

func ThrowException() {
	if r := recover(); r != nil {
		switch reflect.TypeOf(r).String() {
		case "*runtime.TypeAssertionError":
			println(r.(*runtime.TypeAssertionError).Error())
			break
		case "string":
			println(r)
			break
		}
	}
}
