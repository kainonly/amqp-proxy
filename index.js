const env = require('dotenv').config().parsed;
const amqp = require('amqplib');

amqp.connect(env.amqp_uri).then(async (connect) => {
    try {
        const channel = await connect.createChannel();
        channel.assertExchange(env.exchange, 'topic', {durable: true});
        channel.assertQueue(env.queue, {durable: true});
        channel.bindQueue(env.queue, env.exchange);
        channel.consume(env.queue, async (msg) => {
            console.log(msg.content.toString());
        }, {noAck: false});
    } catch (e) {
        console.log(e)
    }
});

