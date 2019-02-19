package main

import (
	"github.com/go-ini/ini"
	"github.com/kainonly/collection-service/src/common"
)

func main() {
	common.Config = new(common.Cogs)

	if err := ini.MapTo(common.Config, "./cogs.ini"); err != nil {
		panic(err.Error())
	}

	println(common.Config.Rabbitmq.Port)
	println(common.Config.Mongodb.Port)
	println(common.Config.Logs.SystemDatabase)

}
