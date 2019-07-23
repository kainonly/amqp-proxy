# AMQP Logging Service

#### 安装

设置 `.env`

```ini
amqp_uri = amqp://kain:123456@imac.com
mongo_uri = mongodb://root:123456@imac.com:27017
database = center
exchange = logging.service
queue = collection
```

#### 默认数据库

白名单 `whitelist` 集合

```json
{
  "appid": "appid",
  "namespace": ["system"]
}
```

#### 队列结构

- `appid` 应用ID
- `namespace` 命名空间
- `data` 收集数据

