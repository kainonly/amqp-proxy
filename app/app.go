package app

import (
	"amqp-proxy/app/controller"
	"amqp-proxy/app/logging"
	"amqp-proxy/app/session"
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
	logger := logging.NewLogging(option.Transfer.Listen)
	ns, err := session.NewSession(option.Amqp, logger, &option.Transfer.Pipe)
	if err != nil {
		return
	}
	pb.RegisterRouterServer(
		server,
		controller.New(ns),
	)
	server.Serve(listen)
	return
}
