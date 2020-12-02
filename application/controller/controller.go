package controller

import (
	pb "amqp-proxy/api"
	"amqp-proxy/application/common"
)

type controller struct {
	pb.UnimplementedAPIServer
	*common.Dependency
}

func New(dep *common.Dependency) *controller {
	c := new(controller)
	c.Dependency = dep
	return c
}
