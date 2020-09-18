package logging

import (
	pb "amqp-proxy/transfer"
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
)

func (c *Logging) Push(pipe string, value interface{}) (err error) {
	var data []byte
	data, err = jsoniter.Marshal(value)
	if err != nil {
		return
	}
	response, err := c.transfer.Push(context.Background(), &pb.PushParameter{
		Identity: pipe,
		Data:     data,
	})
	if err != nil {
		return
	}
	if response.Error != 0 {
		return errors.New(response.Msg)
	}
	return
}
