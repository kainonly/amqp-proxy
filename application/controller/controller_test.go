package controller

import (
	pb "amqp-proxy/api"
	"amqp-proxy/bootstrap"
	"amqp-proxy/config"
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"testing"
)

var client pb.APIClient

func TestMain(m *testing.M) {
	os.Chdir("../../")
	var err error
	var cfg *config.Config
	if cfg, err = bootstrap.LoadConfiguration(); err != nil {
		log.Fatalln(err)
	}
	var conn *grpc.ClientConn
	if conn, err = grpc.Dial(cfg.Listen, grpc.WithInsecure()); err != nil {
		log.Fatalln(err)
	}
	client = pb.NewAPIClient(conn)
	os.Exit(m.Run())
}

func TestController_Publish(t *testing.T) {
	_, err := client.Publish(context.Background(), &pb.Option{
		Exchange:    "debug.proxy",
		Key:         "",
		Mandatory:   false,
		Immediate:   false,
		ContentType: "application/json",
		Body:        []byte(`{"name":"kain"}`),
	})
	if err != nil {
		t.Fatal(err)
	}
}

var receipt string

func TestController_Get(t *testing.T) {
	response, err := client.Get(context.Background(), &pb.Queue{
		Queue: "debug.proxy",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.Body)
	receipt = response.Receipt
}

func TestController_Ack(t *testing.T) {
	_, err := client.Ack(context.Background(), &pb.Receipt{
		Queue:   "debug.proxy",
		Receipt: receipt,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestController_Nack(t *testing.T) {
	_, err := client.Publish(context.Background(), &pb.Option{
		Exchange:    "debug.proxy",
		Key:         "",
		Mandatory:   false,
		Immediate:   false,
		ContentType: "application/json",
		Body:        []byte(`{"name":"nack"}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	// Get Message
	response, err := client.Get(context.Background(), &pb.Queue{
		Queue: "debug.proxy",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response.Body)
	receipt = response.Receipt
	// Nack
	_, err = client.Nack(context.Background(), &pb.Receipt{
		Queue:   "debug.proxy",
		Receipt: receipt,
	})
	if err != nil {
		t.Fatal(err)
	}
}
