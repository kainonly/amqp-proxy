package controller

import (
	pb "amqp-proxy/api"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (c *controller) Ack(_ context.Context, receipt *pb.Receipt) (*empty.Empty, error) {
	if err := c.Session.Ack(receipt.Queue, receipt.Receipt); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
