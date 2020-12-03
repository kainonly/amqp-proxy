package transfer

import (
	pb "amqp-proxy/transfer"
	"context"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
)

type Transfer struct {
	on     bool
	client pb.APIClient
}

func New(listen string) (c *Transfer, err error) {
	c = new(Transfer)
	if listen == "" {
		c.on = false
		return
	}
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(listen, grpc.WithInsecure()); err != nil {
		return
	}
	c.on = true
	c.client = pb.NewAPIClient(conn)
	return
}

func (c *Transfer) Push(pipe string, value interface{}) (err error) {
	if !c.on {
		return
	}
	var data []byte
	if data, err = jsoniter.Marshal(value); err != nil {
		return
	}
	if _, err = c.client.Push(context.Background(), &pb.Body{
		Id:      pipe,
		Content: data,
	}); err != nil {
		return
	}
	return
}
