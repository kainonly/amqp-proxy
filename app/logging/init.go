package logging

import (
	pb "amqp-proxy/transfer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Logging struct {
	transfer pb.RouterClient
}

func NewLogging(listen string) *Logging {
	c := new(Logging)
	conn, err := grpc.Dial(listen, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln(err)
	}
	c.transfer = pb.NewRouterClient(conn)
	return c
}
