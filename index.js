const env = require('dotenv').config().parsed;
const amqp = require('amqplib');
const joi = require('joi');
const {MongoClient} = require('mongodb');
const client = new MongoClient(env.mongo_uri, {useNewUrlParser: true});

amqp.connect(env.amqp_uri).then(async (connect) => {
    try {
        const channel = await connect.createChannel();
        channel.assertExchange(env.exchange, 'topic', {durable: true});
        channel.assertQueue(env.queue, {durable: true});
        channel.bindQueue(env.queue, env.exchange);
        channel.consume(env.queue, async (msg) => {
            try {
                const param = JSON.parse(msg.content.toString());
                const validate = joi.validate(param, joi.object({
                    namespace: joi.string().required(),
                    data: joi.required()
                }));
                if (validate.error !== null) {
                    channel.ack(msg);
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
    console.log(err);
});
