const env = require('dotenv').config().parsed;
const amqp = require('amqplib');
const joi = require('joi');
const Redis = require('ioredis');
const {MongoClient} = require('mongodb');
const client = new MongoClient(env.mongo_uri, {useNewUrlParser: true});
const redis = new Redis(env.redis_uri);
const {Subject} = require('rxjs');

const waitInsert = new Subject();
amqp.connect(env.amqp_uri).then(async (connect) => {
    try {
        const channel = await connect.createChannel();
        channel.assertExchange(env.exchange, 'direct', {durable: true});
        channel.assertQueue(env.queue, {durable: true});
        channel.bindQueue(env.queue, env.exchange);
        channel.consume(env.queue, async (msg) => {
            try {
                const param = JSON.parse(msg.content.toString());
                const validate = joi.validate(param, joi.object({
                    appid: joi.string().required(),
                    namespace: joi.string().required(),
                    data: joi.required()
                }));
                if (validate.error !== null) {
                    channel.ack(msg);
                    return;
                }
                const allowed = await redis.sismember('appid:' + param.appid, param.namespace);
                if (!allowed) {
                    channel.ack(msg);
                } else {
                    waitInsert.next(Object.assign(param, {
                        channel,
                        msg
                    }));
                }
            } catch (e) {
                channel.ack(msg);
            }
        }, {noAck: false});
    } catch (e) {
        console.log(e)
    }
});

client.connect((err) => {
    if (err) return;
    waitInsert.subscribe(async ({appid, namespace, data, channel, msg}) => {
        const database = await redis.hget('daq', appid);
        const db = client.db(database);
        const collection = db.collection(namespace);
        collection.insertOne(data, (err) => {
            if (!err) {
                channel.ack(msg);
            }
        });
    });
});
