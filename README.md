# collection-service

RabbitMQ+MongoDB 简易日志收集服务

#### 配置

创建文件 `cogs.ini` 进行配置

```ini
[rabbitmq]
hostname = developer.com
port = 5672
username = awesome
password = 12345678
vhost = /

[mongodb]
hostname = developer.com
port = 27017
username = root
password = 12345678

[collection]
database = collection_service # 服务数据库名称
exchange = collect # 公共收集交换器
queue = collect # 公共收集队列
system_exchange = log.system # 系统日志收集交换器
system_queue = log.system # 系统日志收集队列
```

初始化服务数据库 `collection_service`, 该数据库下存在 `authorization` 公共验证与 `whitelist` 系统日志收集白名单集合

#### 定义集合

开启服务前首要手动配置数据库集合, 也可以按照该集合结构扩展并对接管理后台

##### - authorization

```json
{
    "app":"collection_developer",
    "appid":"ehVxEb4jjt5HwLkl",
    "secret":"py2GmZOKjSETlcS58wZSddMU9yoByGUL"
}
```

- **app** 所属数据库名称
- **appid** 应用ID
- **secret** 使用密钥

##### - whitelist

```json
{
    "domain":"api.developer.com"
}
```

- **domain** 允许进行系统日志收集的域名

