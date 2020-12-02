package controller

import (
	pb "amqp-proxy/api"
	"context"
)

func (c *controller) Get(_ context.Context, queue *pb.Queue) (content *pb.Content, err error) {
	var receipt string
	var body []byte
	if receipt, body, err = c.Session.Get(queue.Queue); err != nil {
		return
	}
	content = &pb.Content{
		Receipt: receipt,
		Body:    body,
	}
	return
}
