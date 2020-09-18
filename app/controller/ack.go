package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Ack(ctx context.Context, param *pb.AckParameter) (*pb.Response, error) {
	return nil, nil
}
