package common

import (
	"amqp-proxy/application/service/session"
	"amqp-proxy/application/service/transfer"
	"amqp-proxy/config"
	"go.uber.org/fx"
)

type Dependency struct {
	fx.In

	Config   *config.Config
	Session  *session.Session
	Transfer *transfer.Transfer
}
