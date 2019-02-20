package main

import (
	"github.com/go-ini/ini"
	"github.com/kainonly/collection-service/src/collection"
	"github.com/kainonly/collection-service/src/common"
	"github.com/kainonly/collection-service/src/facade"
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

	if err = config.RegisteredMongo(); err != nil {
		panic(err.Error())
	}

	// recover print
	defer facade.AMQPConnection.Close()
	defer facade.AMQPChannel.Close()
	defer facade.Cancel()
	defer func() {
		if r := recover(); r != nil {
			println(r.(string))
		}
	}()

	facade.WG.Add(2)
	// collection system log
	collection.NewSystem(
		config.SystemDatabase,
		config.SystemExchange,
		config.SystemQueue,
	).Subscribe()

	// collection information
	collection.NewStatistics(
		config.StatisticsExchange,
		config.StatisticsQueue,
	).Subscribe()

	facade.WG.Wait()
}
