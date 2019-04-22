const env = require('dotenv').config().parsed;
const amqp = require('amqplib');
const joi = require('joi');
const {MongoClient} = require('mongodb');
const client = new MongoClient(env.mongo_uri, {useNewUrlParser: true});
const {AsyncSubject} = require('rxjs');

const mongo = new AsyncSubject();
client.connect((err) => {
    if (err) return;
    mongo.next(client);
    mongo.complete();
});

amqp.connect(env.amqp_uri).then(async (connect) => {
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

            mongo.subscribe(async (client) => {
                const center = client.db(env.database);
                const data = await center.collection('whitelist').findOne({
                    appid: param.appid
                });
                if (data.namespace.indexOf(param.namespace) === -1) {
                    channel.ack(msg);
                } else {
                    const db = client.db(param.appid);
                    const collection = db.collection(param.namespace);
                    collection.insertOne(param.data, (err) => {
                        if (!err) channel.ack(msg);
                    });
                }
            });
        } catch (e) {
            channel.ack(msg);
        }
    }, {noAck: false});
});


