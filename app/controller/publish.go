package controller

import (
	"amqp-proxy/app/types"
	pb "amqp-proxy/router"
	"context"
)

func (c *controller) Publish(ctx context.Context, param *pb.PublishParameter) (*pb.Response, error) {
	err := c.session.Publish(&types.PublishOption{
		Exchange:    param.Exchange,
		Key:         param.Key,
		Mandatory:   param.Mandatory,
		Immediate:   param.Immediate,
		ContentType: param.ContentType,
		Body:        param.Body,
	})
	if err != nil {
		return c.response(err)
	}
	return c.response(nil)
}
