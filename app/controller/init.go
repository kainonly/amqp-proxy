package controller

import (
	pb "amqp-proxy/router"
)

type controller struct {
	pb.UnimplementedRouterServer
}

func New() *controller {
	c := new(controller)
	return c
}

func (c *controller) response(err error) (*pb.Response, error) {
	if err != nil {
		return &pb.Response{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	} else {
		return &pb.Response{
			Error: 0,
			Msg:   "ok",
		}, nil
	}
}
