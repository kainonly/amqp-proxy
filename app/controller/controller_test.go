package controller

import (
	"amqp-proxy/app/types"
	pb "amqp-proxy/router"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"
)

var client pb.RouterClient

func TestMain(m *testing.M) {
	os.Chdir("../..")
	if _, err := os.Stat("./config/config.yml"); os.IsNotExist(err) {
		logrus.Fatalln("The service configuration file does not exist")
	}
	cfgByte, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		logrus.Fatalln("Failed to read service configuration file", err)
	}
	config := types.Config{}
	err = yaml.Unmarshal(cfgByte, &config)
	if err != nil {
		logrus.Fatalln("Service configuration file parsing failed", err)
	}
	grpcConn, err := grpc.Dial(config.Listen, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln(err)
	}
	client = pb.NewRouterClient(grpcConn)
	os.Exit(m.Run())
}

func TestController_Publish(t *testing.T) {
	response, err := client.Publish(context.Background(), &pb.PublishParameter{
		Exchange:    "test",
		Key:         "",
		Mandatory:   false,
		Immediate:   false,
		ContentType: "application/json",
		Body:        []byte(`{"name":"kain"}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
	response, err = client.Publish(context.Background(), &pb.PublishParameter{
		Exchange:    "test",
		Key:         "",
		Mandatory:   false,
		Immediate:   false,
		ContentType: "application/json",
		Body:        []byte(`{"name":"vvv"}`),
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
}

var receipt1 string
var receipt2 string

func TestController_Get(t *testing.T) {
	response, err := client.Get(context.Background(), &pb.GetParameter{
		Queue: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Data)
		receipt1 = response.Data.Receipt
	}
	response, err = client.Get(context.Background(), &pb.GetParameter{
		Queue: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Data)
		receipt2 = response.Data.Receipt
	}
}

func TestController_Ack(t *testing.T) {
	response, err := client.Ack(context.Background(), &pb.AckParameter{
		Queue:   "test",
		Receipt: receipt1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
}

func TestController_Nack(t *testing.T) {
	response, err := client.Nack(context.Background(), &pb.NackParameter{
		Queue:   "test",
		Receipt: receipt2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if response.Error != 0 {
		t.Error(response.Msg)
	} else {
		t.Log(response.Msg)
	}
}

func BenchmarkController(b *testing.B) {
	for i := 0; i < 1000; i++ {
		response, err := client.Publish(context.Background(), &pb.PublishParameter{
			Exchange:    "test",
			Key:         "",
			Mandatory:   false,
			Immediate:   false,
			ContentType: "application/json",
			Body:        []byte(`{"name":"test"}`),
		})
		if err != nil {
			b.Fatal(err)
		}
		if response.Error != 0 {
			b.Error(response.Msg)
		}
	}
}

func BenchmarkController_GetAndAck(b *testing.B) {
	for i := 0; i < 1000; i++ {
		response1, err := client.Get(context.Background(), &pb.GetParameter{
			Queue: "test",
		})
		if err != nil {
			b.Fatal(err)
		}
		if response1.Error != 0 {
			b.Error(response1.Msg)
		}
		response2, err := client.Ack(context.Background(), &pb.AckParameter{
			Queue:   "test",
			Receipt: response1.Data.Receipt,
		})
		if err != nil {
			b.Fatal(err)
		}
		if response2.Error != 0 {
			b.Error(response2.Msg)
		}
	}
}

func BenchmarkController_Mock(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 5000; i++ {
			response, err := client.Publish(context.Background(), &pb.PublishParameter{
				Exchange:    "test",
				Key:         "",
				Mandatory:   false,
				Immediate:   false,
				ContentType: "application/json",
				Body:        []byte(`{"name":"test"}`),
			})
			if err != nil {
				b.Fatal(err)
			}
			if response.Error != 0 {
				b.Error(response.Msg)
			} else {
				println("send message: (", i, ")")
			}
		}
		wg.Done()
	}()
	time.Sleep(time.Second)
	go func() {
		for i := 0; i < 5000; i++ {
			response1, err := client.Get(context.Background(), &pb.GetParameter{
				Queue: "test",
			})
			if err != nil {
				b.Fatal(err)
			}
			if response1.Error != 0 {
				b.Error(response1.Msg)
			} else {
				println("get message: (", i, ")")
			}
			if i%2 == 0 {
				response2, err := client.Ack(context.Background(), &pb.AckParameter{
					Queue:   "test",
					Receipt: response1.Data.Receipt,
				})
				if err != nil {
					b.Fatal(err)
				}
				if response2.Error != 0 {
					b.Error(response2.Msg)
				} else {
					println("ack message: (", i, ")")
				}
			} else {
				response2, err := client.Nack(context.Background(), &pb.NackParameter{
					Queue:   "test",
					Receipt: response1.Data.Receipt,
				})
				if err != nil {
					b.Fatal(err)
				}
				if response2.Error != 0 {
					b.Error(response2.Msg)
				} else {
					println("nack message: (", i, ")")
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
