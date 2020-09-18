package app

import (
	"amqp-proxy/app/controller"
	"amqp-proxy/app/types"
	pb "amqp-proxy/router"
	"google.golang.org/grpc"
	"net"
	"net/http"
	_ "net/http/pprof"
)

func Application(option *types.Config) (err error) {
	// Turn on debugging
	if option.Debug {
		go func() {
			http.ListenAndServe(":6060", nil)
		}()
	}
	// Start microservice
	listen, err := net.Listen("tcp", option.Listen)
	if err != nil {
		return
	}
	server := grpc.NewServer()
	pb.RegisterRouterServer(
		server,
		controller.New(),
	)
	server.Serve(listen)
	return
}
