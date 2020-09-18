package controller

import (
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Get(ctx context.Context, param *pb.GetParameter) (*pb.GetResponse, error) {
	receipt, body, err := c.session.Get(param.Queue)
	if err != nil {
		return &pb.GetResponse{
			Error: 1,
			Msg:   err.Error(),
		}, nil
	}
	return &pb.GetResponse{
		Error: 0,
		Msg:   "ok",
		Data: &pb.Data{
			Receipt: receipt,
			Body:    body,
		},
	}, nil
}
