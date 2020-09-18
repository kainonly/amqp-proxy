package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Ack(ctx context.Context, param *pb.AckParameter) (*pb.Response, error) {
	err := c.session.Ack(param.Queue, param.Receipt)
	if err != nil {
		return c.response(err)
	}
	return c.response(nil)
}
