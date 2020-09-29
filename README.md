# AMQP Proxy

Use grpc proxy to call AMQP operations

[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/codexset/amqp-proxy?style=flat-square)](https://github.com/codexset/amqp-proxy)
[![Github Actions](https://img.shields.io/github/workflow/status/codexset/amqp-proxy/release?style=flat-square)](https://github.com/codexset/amqp-proxy/actions)
[![Image Size](https://img.shields.io/docker/image-size/kainonly/amqp-proxy?style=flat-square)](https://hub.docker.com/r/kainonly/amqp-proxy)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/amqp-proxy.svg?style=flat-square)](https://hub.docker.com/r/kainonly/amqp-proxy)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/codexset/amqp-proxy/master/LICENSE)

## Setup

Example using docker compose

```yaml
version: "3.8"
services: 
  subscriber:
    image: kainonly/amqp-proxy
    restart: always
    volumes:
      - ./amqp-proxy:/app/config
    ports:
      - 6000:6000
```

## Configuration

For configuration, please refer to `config/config.example.yml`

- **debug** `string` Start debugging, ie `net/http/pprof`, access address is`http://localhost:6060`
- **listen** `string` Microservice listening address
- **amqp** `string` E.g `amqp://guest:guest@localhost:5672/`
- **transfer** `object` [elastic-transfer](https://github.com/codexset/elastic-transfer) service
  - **listen** `string` host
  - **pipe** `object`
    - **publish** `string` for `publish`
    - **message** `string` for `get` `ack` `nack`

## Service

The service is based on gRPC and you can view `router/router.proto`

```proto
syntax = "proto3";
package amqp.proxy;
service Router {
  rpc Publish (PublishParameter) returns (Response){
  }

  rpc Get (GetParameter) returns (GetResponse) {
  }

  rpc Ack (AckParameter) returns (Response) {
  }

  rpc Nack (NackParameter) returns (Response){
  }
}

message Response {
  uint32 error = 1;
  string msg = 2;
}

message PublishParameter{
  string exchange = 1;
  string key = 2;
  bool mandatory = 3;
  bool immediate = 4;
  string contentType = 5;
  bytes body = 6;
}

message GetParameter {
  string queue = 1;
}

message GetResponse {
  uint32 error = 1;
  string msg = 2;
  Data data = 3;
}

message Data {
  string receipt = 1;
  bytes body = 2;
}

message AckParameter {
  string queue = 1;
  string receipt = 2;
}

message NackParameter {
  string queue = 1;
  string receipt = 2;
}
```

## rpc Publish (PublishParameter) returns (Response) {}

- PublishParameter
  - **exchange** `string` exchange name
  - **key** `string` routing key
  - **mandatory** `bool`
  - **immediate** `bool`
  - **contentType** `string` text/plain or application/json 
  - **body** `bytes` publish payload
- Response
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback

```golang
client.Publish(context.Background(), &pb.PublishParameter{
    Exchange:    "proxy",
    Key:         "",
    Mandatory:   false,
    Immediate:   false,
    ContentType: "application/json",
    Body:        []byte(`{"name":"kain"}`),
})
```

## rpc Get (GetParameter) returns (GetResponse) {}

- GetParameter
  - **queue** `string` queue name
- GetResponse
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback
  - **data** `Data`
    - **receipt** `string` consumption receipt
    - **body** `bytes` get payload

```golang
client.Get(context.Background(), &pb.GetParameter{
    Queue: "proxy",
})
```

## rpc Ack (AckParameter) returns (Response) {}

- AckParameter
  - **queue** `string` queue name
  - **receipt** `string` consumption receipt
- Response
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback

```golang
client.Ack(context.Background(), &pb.AckParameter{
    Queue:   "proxy",
    Receipt: receipt,
})
```

## rpc Nack (NackParameter) returns (Response) {}

- NackParameter
  - **queue** `string` queue name
  - **receipt** `string` consumption receipt
- Response
  - **error** `uint32` error code, `0` is normal
  - **msg** `string` error feedback

```golang
client.Nack(context.Background(), &pb.NackParameter{
    Queue:   "proxy",
    Receipt: receipt,
})
```