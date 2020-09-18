package controller

import (
	"amqp-proxy/app/session"
	pb "amqp-proxy/router"
)

type controller struct {
	pb.UnimplementedRouterServer
	session *session.Session
}

func New(session *session.Session) *controller {
	c := new(controller)
	c.session = session
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
