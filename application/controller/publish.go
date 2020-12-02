package controller

import (
	pb "amqp-proxy/api"
	"amqp-proxy/application/service/session"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (c *controller) Publish(_ context.Context, option *pb.Option) (*empty.Empty, error) {
	if err := c.Session.Publish(session.PublishOption{
		Exchange:    option.Exchange,
		Key:         option.Key,
		Mandatory:   option.Mandatory,
		Immediate:   option.Immediate,
		ContentType: option.ContentType,
		Body:        option.Body,
	}); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
