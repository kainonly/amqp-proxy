package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Nack(ctx context.Context, param *pb.NackParameter) (*pb.Response, error) {
	err := c.session.Nack(param.Queue, param.Receipt)
	if err != nil {
		return c.response(err)
	}
	return c.response(nil)
}
