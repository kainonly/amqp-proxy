# AMQP Proxy

Use grpc proxy to call AMQP operations

[![Github Actions](https://img.shields.io/github/workflow/status/kain-lab/amqp-proxy/release?style=flat-square)](https://github.com/kain-lab/amqp-proxy/actions)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kain-lab/amqp-proxy?style=flat-square)](https://github.com/kain-lab/amqp-proxy)
[![Image Size](https://img.shields.io/docker/image-size/kainonly/amqp-proxy?style=flat-square)](https://hub.docker.com/r/kainonly/amqp-proxy)
[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/amqp-proxy.svg?style=flat-square)](https://hub.docker.com/r/kainonly/amqp-proxy)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kain-lab/amqp-proxy/master/LICENSE)

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
      - 8080:8080
```

## Configuration

For configuration, please refer to `config/config.example.yml`

- **debug** `string` Start debugging, ie `net/http/pprof`, access address is`http://localhost:6060`
- **listen** `string` grpc server listening address
- **gateway** `string` API gateway server listening address
- **amqp** `string` E.g `amqp://guest:guest@localhost:5672/`
- **transfer** `object` [elastic-transfer](https://github.com/kain-lab/elastic-transfer) service
  - **listen** `string` host
  - **pipe** `object`
    - **publish** `string` for `publish`
    - **message** `string` for `get` `ack` `nack`

## Service

The service is based on gRPC to view `api/api.proto`

```proto
syntax = "proto3";
package amqp.proxy;
option go_package = "amqp-proxy/gen/go/amqp/proxy";
import "google/protobuf/empty.proto";
import "google/api/annotations.proto";

service API {
  rpc Publish (Option) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/publish",
      body: "*"
    };
  }
  rpc Get (Queue) returns (Content) {
    option (google.api.http) = {
      get: "/get",
    };
  }
  rpc Ack (Receipt) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/ack",
      body: "*"
    };
  }
  rpc Nack (Receipt) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/nack",
      body: "*"
    };
  }
}

message Option {
  string exchange = 1;
  string key = 2;
  bool mandatory = 3;
  bool immediate = 4;
  string contentType = 5;
  bytes body = 6;
}

message Queue {
  string queue = 1;
}

message Content {
  string receipt = 1;
  bytes body = 2;
}

message Receipt {
  string queue = 1;
  string receipt = 2;
}
```

## Publish (Option) returns (google.protobuf.Empty)

Publish messages to the exchange

### RPC

- **Option**
  - **exchange** `string` exchange name
  - **key** `string` routing key
  - **mandatory** `bool`
  - **immediate** `bool`
  - **contentType** `string` text/plain or application/json 
  - **body** `bytes` publish payload

```golang
client := pb.NewAPIClient(conn)
_, err := client.Publish(context.Background(), &pb.Option{
  Exchange:    "proxy.debug",
  Key:         "",
  Mandatory:   false,
  Immediate:   false,
  ContentType: "application/json",
  Body:        []byte(`{"name":"kain"}`),
})
```

### API Gateway

- **POST** `/publish`

```http
POST /publish HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "exchange": "proxy.debug",
    "key": "",
    "mandatory": false,
    "immediate": false,
    "content_type": "application/json",
    "body": "eyJuYW1lIjoiYXBpIn0="
}
```

## Get (Queue) returns (Content)

Get messages from the queue

### RPC

- **Queue**
  - **queue** `string` queue name
- **Content**
  - **receipt** `string` consumption receipt
  - **body** `bytes` get payload

```golang
client := pb.NewAPIClient(conn)
response, err := client.Get(context.Background(), &pb.Queue{
  Queue: "proxy.debug",
})
```

### API Gateway

- **GET** `/get`

```http
GET /get?queue=proxy.debug HTTP/1.1
Host: localhost:8080
```

## Ack (Receipt) returns (google.protobuf.Empty)

Message acknowledgment

### RPC

- **Receipt**
  - **queue** `string` queue name
  - **receipt** `string` consumption receipt

```golang
client := pb.NewAPIClient(conn)
_, err := client.Ack(context.Background(), &pb.Receipt{
  Queue:   "proxy.debug",
  Receipt: receipt,
})
```

### API Gateway

- **POST** `/ack`

```http
POST /ack HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "queue": "proxy.debug",
    "receipt": "456a55fd-dc0f-4073-82ba-74c9c16c4961"
}
```

## Nack (Receipt) returns (google.protobuf.Empty)

Message unacknowledged

### RPC

- **Receipt**
  - **queue** `string` queue name
  - **receipt** `string` consumption receipt

```golang
client := pb.NewAPIClient(conn)
_, err = client.Nack(context.Background(), &pb.Receipt{
  Queue:   "proxy.debug",
  Receipt: receipt,
})
```

### API Gateway

- **POST** `/nack`

```http
POST /nack HTTP/1.1
Host: localhost:8080
Content-Type: application/json

{
    "queue": "proxy.debug",
    "receipt": "063207ab-603a-4f1c-acda-97ffe089fb52"
}
```