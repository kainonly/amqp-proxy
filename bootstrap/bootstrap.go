package bootstrap

import (
	pb "amqp-proxy/api"
	"amqp-proxy/application/service/session"
	"amqp-proxy/application/service/transfer"
	"amqp-proxy/config"
	"context"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var (
	LoadConfigurationNotExists = errors.New("the configuration file does not exist")
)

// Load application configuration
// reference config.example.yml
func LoadConfiguration() (cfg *config.Config, err error) {
	if _, err = os.Stat("./config/autoload"); os.IsNotExist(err) {
		os.Mkdir("./config/autoload", os.ModeDir)
	}
	if _, err = os.Stat("./config/config.yml"); os.IsNotExist(err) {
		err = LoadConfigurationNotExists
		return
	}
	var bs []byte
	if bs, err = ioutil.ReadFile("./config/config.yml"); err != nil {
		return
	}
	if err = yaml.Unmarshal(bs, &cfg); err != nil {
		return
	}
	return
}

// Initialize transfer client
func InitializeTransfer(cfg *config.Config) (*transfer.Transfer, error) {
	return transfer.New(cfg.Transfer.Listen)
}

// Initialize amqp session management
func InitializeSession(dep session.Dependency) (*session.Session, error) {
	return session.New(&dep)
}

// Set up api gateway for grpc
func ApiGateway(cfg *config.Config) (err error) {
	if cfg.Debug != "" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	ctx := context.Background()
	mux := runtime.NewServeMux()
	endpoint := *flag.String("grpc-server-endpoint", cfg.Listen, "gRPC server endpoint")
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err = pb.RegisterAPIHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return
	}
	serve := gin.New()
	serve.Any("/*any", gin.WrapF(mux.ServeHTTP))
	serve.Run(cfg.Gateway)
	return
}
