package main

import (
	"github.com/go-ini/ini"
	"github.com/kainonly/collection-service/common"
	"github.com/kainonly/collection-service/controller"
	"github.com/kainonly/collection-service/facade"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var err error

func main() {
	config := new(common.Cogs)
	// load env
	if err = ini.MapTo(config, "cogs.ini"); err != nil {
		panic(err.Error())
	}

	if config.ValidateArgs() {
		panic("please set cogs.ini!")
	}

	if err = config.RegisteredAMQP(); err != nil {
		panic(err.Error())
	}

	defer facade.AMQPConnection.Close()
	defer facade.AMQPChannel.Close()

	facade.Db = make(map[string]*mongo.Database)
	if err = config.RegisteredMongo(); err != nil {
		panic(err.Error())
	}

	facade.WG.Add(2)

	// collection information
	controller.NewStatistics(
		config.Collection.Database,
		config.Collection.Exchange,
		config.Collection.Queue,
	).Run()

	// collection system log
	controller.NewSystem(
		config.Collection.SystemDatabase,
		config.Collection.SystemExchange,
		config.Collection.SystemQueue,
	).Run()

	facade.WG.Wait()
}
