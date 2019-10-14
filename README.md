# Elastic AMQP Stash

Elasticsearch Queue Pipeline Data Collector

[![Docker Pulls](https://img.shields.io/docker/pulls/kainonly/elastic-amqp-stash.svg?style=flat-square)](https://hub.docker.com/r/kainonly/elastic-amqp-stash)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/kainonly/elastic-amqp-stash.svg?style=flat-square)](https://hub.docker.com/r/kainonly/elastic-amqp-stash)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/kainonly/elastic-amqp-stash.svg?style=flat-square)](https://hub.docker.com/r/kainonly/elastic-amqp-stash)
[![TypeScript](https://img.shields.io/badge/%3C%2F%3E-TypeScript-blue.svg?style=flat-square)](https://github.com/kainonly/elastic-amqp-stash)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kainonly/elastic-amqp-stash/master/LICENSE)

```shell
docker pull kainonly/elastic-amqp-stash
```

## Docker Compose

example

```yml
version: '3.7'
services:
  stash:
    image: kainonly/elastic-amqp-stash
    restart: always
    volumes: 
      - ./stash/config.json:/app/data/config.json
```

## How to use

set config.json

- **amqp** amqp connect options, detail [amqplib docs](http://www.squaremobius.net/amqp.node/channel_api.html)
- **elastic** elasticsearch connect options, detail [client configuration](https://www.elastic.co/guide/en/elasticsearch/client/javascript-api/current/client-configuration.html)
- **rule** Use JSON Schema to verify the body that the elastic index needs to write. If the corresponding index rule is not set, it will not be written.

```json
{
  "amqp": {
    "hostname": "localhost",
    "username": "guest",
    "password": "guest"
  },
  "elastic": {
    "node": "http://localhost:9200"
  },
  "rule": {
    "system": {
        "type": "object",
        "required": [
            "appid",
            "username",
            "role",
            "symbol",
            "request",
            "method",
            "query",
            "body",
            "ip",
            "user_agent",
            "create_time"
        ],
        "properties": {
            "appid": {
                "type": "string"
            },
            "username": {
                "type": "string"
            },
            "role": {
                "type": [
                    "string",
                    "array"
                ]
            },
            "symbol": {
                "type": "object"
            },
            "request": {
                "type": "string"
            },
            "method": {
                "type": "string"
            },
            "query": {
                "type": "array"
            },
            "body": {
                "type": "array"
            },
            "ip": {
                "type": "string",
                "format": "ipv4"
            },
            "user_agent": {
                "type": "string"
            },
            "create_time": {
                "type": "number"
            }
        }
    }
  }
}
```

## send and stash

default amqp setting

- **exchange** `elastic.stash.exchange`
- **queue** `elastic.stash.basic`

queue message needs to be posted to `elastic.stash.exchange`

```json
{
    "index": "system",
    "body": {
        "appid": "xx-1",
        "username": "kain",
        "role": ["*"],
        "symbol": {
          "staff": true
        },
        "request": "https://api.developer.com/test/index",
        "method": "POST",
        "query": {},
        "body": {"name": "kain"},
        "ip": "127.0.0.1",
        "user_agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) snap Chromium/75.0.3770.142 Chrome/75.0.3770.142 Safari/537.36",
        "create_time": 1563932873
    }
}
```
