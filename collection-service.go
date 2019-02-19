package main

import (
	"github.com/go-ini/ini"
	"github.com/kainonly/collection-service/src/common"
	"github.com/kainonly/collection-service/src/logs"
)

var err error

func main() {
	config := new(common.Cogs)
	// load env
	if err = ini.MapTo(config, "cogs.ini"); err != nil {
		panic(err.Error())
	}

	if config.CheckArgs() {
		panic("please set cogs.ini!")
	}

	if err = config.RegisteredAMQP(); err != nil {
		panic(err.Error())
	}

	if err = config.RegisteredMongo(); err != nil {
		panic(err.Error())
	}

	// running application
	logs.NewSystem(
		config.SystemDatabase,
		config.SystemExchange,
		config.SystemQueue,
	).Subscribe()

	// recover panic
	defer func() {
		if r := recover(); r != nil {
			println("error:", r)
		}
	}()
}
