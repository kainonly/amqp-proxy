package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Get(ctx context.Context, param *pb.GetParameter) (*pb.GetResponse, error) {
	return nil, nil
}
