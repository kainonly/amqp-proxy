package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Publish(ctx context.Context, param *pb.PublishParameter) (*pb.Response, error) {
	return nil, nil
}
