package main

import (
	"amqp-proxy/application"
	"amqp-proxy/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		//fx.NopLogger,
		fx.Provide(
			bootstrap.LoadConfiguration,
			bootstrap.InitializeTransfer,
			bootstrap.InitializeSession,
		),
		fx.Invoke(application.Application),
	).Run()
}
