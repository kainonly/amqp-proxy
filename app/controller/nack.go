package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Nack(ctx context.Context, param *pb.NackParameter) (*pb.Response, error) {
	return nil, nil
}
